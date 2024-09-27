package search

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2/log"

	"github.com/MikelSot/tribal-training-search/model"
)

type Search struct {
	itunes      ItunesService
	chartLyrics ChartLyricsService
}

func New(itunes ItunesService, chartLyrics ChartLyricsService) Search {
	return Search{
		itunes:      itunes,
		chartLyrics: chartLyrics,
	}
}

func (s Search) Search(ctx context.Context, search string) (model.Results, error) {
	itunesResult, err := s.itunes.Search(ctx, search)
	if err != nil {
		return nil, err
	}

	chartLyricsResult, err := s.chartLyrics.Search(ctx, search)
	if err != nil {
		return nil, err
	}

	if itunesResult.IsEmpty() && chartLyricsResult.IsEmpty() {
		log.Warn("search.Search(): no results found")

		return nil, nil
	}

	itunesResultMap := itunesResult.ToMapByTrackId()

	var results model.Results

	for _, itunes := range itunesResultMap {
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

	for _, chartLyrics := range chartLyricsResult.SearchLyricResult {
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

	return results, nil
}
