package search

import (
	"github.com/MikelSot/tribal-training-search/domain/search"
	"github.com/MikelSot/tribal-training-search/model"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	useCase search.UseCase
}

func newHandler(useCase search.UseCase) handler {
	return handler{useCase}
}

func (h handler) Search(ctx *fiber.Ctx) error {
	searchMap := model.SearchMap{
		model.Artist: model.Search{Search: ctx.Query("artist")},
		model.Song:   model.Search{Search: ctx.Query("song")},
		model.Album:  model.Search{Search: ctx.Query("album")},
	}

	result, err := h.useCase.Search(ctx.Context(), searchMap)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}
