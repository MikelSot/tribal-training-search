package search

import (
	"github.com/MikelSot/tribal-training-search/domain/search"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	useCase search.UseCase
}

func newHandler(useCase search.UseCase) handler {
	return handler{useCase}
}

func (h handler) Search(ctx *fiber.Ctx) error {
	param := ctx.Params("search")

	if param == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON("search param is required")
	}

	result, err := h.useCase.Search(ctx.Context(), param)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}
