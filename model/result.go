package model

type Result struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Artist   string  `json:"artist"`
	Duration string  `json:"duration"`
	Album    string  `json:"album"`
	Artwork  string  `json:"artwork"`
	Price    float32 `json:"price"`
	Origin   Origin  `json:"origin"`
}

type Results []Result
