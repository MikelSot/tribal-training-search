package search

import (
	"context"
	"time"

	"github.com/MikelSot/tribal-training-search/model"
)

type UseCase interface {
	Search(ctx context.Context, searches model.SearchMap) (model.Results, error)
}

type ItunesService interface {
	Search(ctx context.Context, search model.Search) (model.ItunesResult, error)
}

type ChartLyricsService interface {
	Search(ctx context.Context, search model.SearchMap) (model.ChartLyricsResult, error)
}

type RedisService interface {
	Set(ctx context.Context, key string, value interface{}, exp time.Duration) error
	Get(ctx context.Context, key string) (string, error)
}
