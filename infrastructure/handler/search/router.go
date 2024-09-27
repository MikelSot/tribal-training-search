package search

import (
	"github.com/MikelSot/tribal-training-search/domain/search"
	chartLyricsService "github.com/MikelSot/tribal-training-search/infrastructure/chartlyrics/search"
	itunesService "github.com/MikelSot/tribal-training-search/infrastructure/itunes/search"
	"github.com/MikelSot/tribal-training-search/model"
	"github.com/gofiber/fiber/v2"
)

const (
	_routePrefix = "/search"
)

func NewRouter(config model.Config) {
	h := buildHandler(config)

	route(config.App, h)
}

func buildHandler(config model.Config) handler {
	itunesSearchService := itunesService.New(config)
	chartLyricsSearchService := chartLyricsService.New(config)

	useCase := search.New(itunesSearchService, chartLyricsSearchService)

	handler := newHandler(useCase)

	return handler
}

func route(app *fiber.App, h handler) {
	api := app.Group(_routePrefix)

	api.Get("/:search", h.Search)
}
