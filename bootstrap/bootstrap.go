package bootstrap

import (
	"context"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"

	"github.com/MikelSot/tribal-training-search/infrastructure/handler"
	"github.com/MikelSot/tribal-training-search/model"
)

func Run() {
	_ = godotenv.Load()

	ctx := context.Background()

	app := newFiber()

	redis := NewRedisClient(ctx)

	handler.InitRouter(model.Config{
		App:            app,
		ItunesUrl:      getItunesRoute(),
		ChartLyricsUrl: getChartsLyricsRoute(),
		Redis:          redis,
	})

	err := app.Listen(getPort())
	if err != nil {
		log.Errorf("Error: %v", err)
		return
	}
}
