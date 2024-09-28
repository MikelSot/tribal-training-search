package model

import (
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	App            *fiber.App
	ItunesUrl      string
	ChartLyricsUrl string
	Redis          *redis.Client
}
