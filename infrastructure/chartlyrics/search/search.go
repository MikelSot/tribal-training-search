package search

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/MikelSot/tribal-training-search/model"
)

const (
	_paramArtist = "artist"
	_paramSong   = "song"
)

type Search struct {
	config model.Config
}

func New(config model.Config) Search {
	return Search{config}
}

func (s Search) Search(ctx context.Context, search string) (model.ChartLyricsResult, error) {
	chartLyricsUrl, err := url.Parse(s.config.ChartLyricsUrl)
	if err != nil {
		return model.ChartLyricsResult{}, fmt.Errorf("search.url.Parse(): %w", err)
	}

	params := url.Values{}
	params.Add(_paramArtist, search)
	params.Add(_paramSong, search)
	chartLyricsUrl.RawQuery = params.Encode()

	body, err := s.doRequest(ctx, chartLyricsUrl.String())
	if err != nil {
		return model.ChartLyricsResult{}, fmt.Errorf("search.doRequest: %w", err)
	}

	var result model.ChartLyricsResult
	if err := xml.Unmarshal(body, &result); err != nil {
		return model.ChartLyricsResult{}, fmt.Errorf("search.json.Unmarshal(): %w", err)
	}

	return result, nil
}

func (s Search) doRequest(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, bytes.NewBuffer([]byte{}))
	if err != nil {
		return nil, fmt.Errorf("search.NewRequestWithContext(): %w", err)
	}

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
