package search

import (
	"context"

	"github.com/MikelSot/tribal-training-search/model"
)

type UseCase interface {
	Search(ctx context.Context, search string) (model.Results, error)
}

type ItunesService interface {
	Search(ctx context.Context, search string) (model.ItunesResult, error)
}

type ChartLyricsService interface {
	Search(ctx context.Context, search string) (model.ChartLyricsResult, error)
}
