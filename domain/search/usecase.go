package search

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2/log"

	"github.com/MikelSot/tribal-training-search/model"
)

type Search struct {
	itunes      ItunesService
	chartLyrics ChartLyricsService
	redis       RedisService
}

func New(itunes ItunesService, chartLyrics ChartLyricsService, redis RedisService) Search {
	return Search{
		itunes:      itunes,
		chartLyrics: chartLyrics,
		redis:       redis,
	}
}

func (s Search) Search(ctx context.Context, searches model.SearchMap) (model.Results, error) {
	cacheKey := fmt.Sprintf("%s%s%s", searches[model.Song].Search, searches[model.Artist].Search, searches[model.Album].Search)

	if len(searches) == 0 {
		return nil, fmt.Errorf("search.Search(): searches is empty")
	}

	results, err := s.getDataFromCache(ctx, cacheKey)
	if err == nil {
		return results, nil
	}
	log.Warn("search.Search(): no results found in cache")

	itunesResult, err := s.getSongsFromItunes(ctx, searches)
	if err != nil {
		return nil, err
	}

	chartLyricsResult, err := s.getSongsFromChartLyrics(ctx, searches)
	if err != nil {
		return nil, err
	}

	if itunesResult.IsEmpty() && chartLyricsResult.IsEmpty() {
		log.Warn("search.Search(): no results found")

		return nil, nil
	}

	itunesResultMap := itunesResult.ToMapByTrackId()

	results = s.makeResults(itunesResultMap, chartLyricsResult)

	if err = s.redis.Set(ctx, cacheKey, results, time.Hour*24); err != nil {
		log.Warn("search.Search(): error saving results in cache")
	}

	return results, nil
}

func (s Search) getSongsFromItunes(ctx context.Context, searches model.SearchMap) (model.ItunesResult, error) {
	resultItunesCh := make(chan model.ItunesResult, len(searches))
	errorsCh := make(chan error, len(searches))

	var wg sync.WaitGroup
	var mx sync.Mutex

	for _, search := range searches {
		wg.Add(1)

		go func(search model.Search, wg *sync.WaitGroup, mx *sync.Mutex) {
			defer wg.Done()

			ctx = context.Background()

			itunesResult, err := s.itunes.Search(ctx, search)
			if err != nil {
				mx.Lock()
				errorsCh <- fmt.Errorf("%w", err)
				mx.Unlock()

				return
			}

			mx.Lock()
			resultItunesCh <- itunesResult
			mx.Unlock()
		}(search, &wg, &mx)
	}

	wg.Wait()
	close(resultItunesCh)
	close(errorsCh)

	var sliceErrors []error
	for err := range errorsCh {
		sliceErrors = append(sliceErrors, err)
	}

	if len(sliceErrors) > 0 {
		return model.ItunesResult{}, fmt.Errorf("search.getSongsFromItunes(): %v", sliceErrors)
	}

	var itunesResult model.ItunesResult
	for result := range resultItunesCh {
		itunesResult.Results = append(itunesResult.Results, result.Results...)
	}

	return itunesResult, nil
}

func (s Search) getSongsFromChartLyrics(ctx context.Context, searches model.SearchMap) (model.ChartLyricsResult, error) {
	song, okSong := searches[model.Song]
	artist, okArtist := searches[model.Artist]
	if (!okSong && !okArtist) || (song.IsSearchEmpty() && artist.IsSearchEmpty()) {
		return model.ChartLyricsResult{}, nil
	}

	if song.IsSearchEmpty() && !artist.IsSearchEmpty() {
		searches[model.Song] = artist
	}

	if !song.IsSearchEmpty() && artist.IsSearchEmpty() {
		searches[model.Artist] = song
	}

	return s.chartLyrics.Search(ctx, searches)
}

func (s Search) makeResults(itunes map[int]model.ItunesResponse, chartLyrics model.ChartLyricsResult) model.Results {
	var results model.Results

	for _, itunes := range itunes {
		result := model.Result{
			Id:       itunes.TrackId,
			Name:     itunes.TrackName,
			Artist:   itunes.ArtistName,
			Duration: strconv.Itoa(itunes.TrackTimeMillis),
			Album:    itunes.CollectionName,
			Artwork:  itunes.ArtworkUrl,
			Price:    itunes.TrackPrice,
			Origin:   model.Itunes,
		}

		results = append(results, result)
	}

	for _, chartLyrics := range chartLyrics.SearchLyricResult {
		if chartLyrics.TrackId == 0 {
			log.Warn("search.Search(): no trackId found in chartLyrics")

			continue
		}

		result := model.Result{
			Id:       chartLyrics.TrackId,
			Name:     chartLyrics.Song,
			Artist:   chartLyrics.Artist,
			Duration: "",
			Album:    "",
			Artwork:  chartLyrics.SongUrl,
			Price:    0,
			Origin:   model.ChartLyrics,
		}

		results = append(results, result)
	}

	return results
}

func (s Search) getDataFromCache(ctx context.Context, key string) (model.Results, error) {
	var results model.Results

	cacheResults, err := s.redis.Get(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("search.getDataFromCache(): %w", err)
	}

	if err = json.Unmarshal([]byte(cacheResults), &results); err != nil {
		return nil, fmt.Errorf("search.getDataFromCache(): %w", err)
	}

	return results, nil
}
