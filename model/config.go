package model

import "github.com/gofiber/fiber/v2"

type Config struct {
	App            *fiber.App
	ItunesUrl      string
	ChartLyricsUrl string
}
