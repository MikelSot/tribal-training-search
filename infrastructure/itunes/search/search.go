package search

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/MikelSot/tribal-training-search/model"
)

const (
	_headerContentType      = "Content-Type"
	_headerContentTypeValue = "application/json"
)

const (
	_paramTerm  = "term"
	_paramMedia = "media"

	_paramEntity = "entity"
	_Media       = "music"
)

type Search struct {
	config model.Config
}

func New(config model.Config) Search {
	return Search{config}
}

func (s Search) Search(ctx context.Context, search model.Search) (model.ItunesResult, error) {
	itunesUrl, err := url.Parse(s.config.ItunesUrl)
	if err != nil {
		return model.ItunesResult{}, fmt.Errorf("search.url.Parse(): %w", err)
	}

	if search.Entity == model.Artist {
		search.Entity = model.MusicArtist
	}

	params := url.Values{}
	params.Add(_paramTerm, search.Search)
	params.Add(_paramEntity, string(search.Entity))
	params.Add(_paramMedia, _Media)
	itunesUrl.RawQuery = params.Encode()

	body, err := s.doRequest(ctx, itunesUrl.String())
	if err != nil {
		return model.ItunesResult{}, fmt.Errorf("search.doRequest: %w", err)
	}

	var result model.ItunesResult
	if err := json.Unmarshal(body, &result); err != nil {
		return model.ItunesResult{}, fmt.Errorf("search.json.Unmarshal(): %w", err)
	}

	return result, nil
}

func (s Search) doRequest(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, bytes.NewBuffer([]byte{}))
	if err != nil {
		return nil, fmt.Errorf("search.NewRequestWithContext(): %w", err)
	}

	req.Header.Add(_headerContentType, _headerContentTypeValue)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client.Do(): %w", err)
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("reading Body: %v", err)
	}

	if res.StatusCode >= http.StatusInternalServerError {
		return nil, fmt.Errorf("Error al obtener las canciones de iTunes")
	}

	return bodyBytes, nil
}
