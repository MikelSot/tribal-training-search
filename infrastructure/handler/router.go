package handler

import (
	"github.com/MikelSot/tribal-training-search/infrastructure/handler/search"
	"github.com/MikelSot/tribal-training-search/model"
)

func InitRouter(config model.Config) {
	// S
	search.NewRouter(config)
}
