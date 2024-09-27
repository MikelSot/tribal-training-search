package bootstrap

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"

	"github.com/MikelSot/tribal-training-search/infrastructure/handler"
	"github.com/MikelSot/tribal-training-search/model"
)

func Run() {
	_ = godotenv.Load()

	app := newFiber()

	handler.InitRouter(model.Config{
		App:            app,
		ItunesUrl:      getItunesRoute(),
		ChartLyricsUrl: getChartsLyricsRoute(),
	})

	err := app.Listen(getPort())
	if err != nil {
		log.Errorf("Error: %v", err)
		return
	}
}
