package model

type ItunesResult struct {
	Results []ItunesResponse `json:"results"`
}

func (i ItunesResult) IsEmpty() bool { return len(i.Results) == 0 }

func (i ItunesResult) ToMapByTrackId() map[int]ItunesResponse {
	m := make(map[int]ItunesResponse)

	for _, result := range i.Results {
		m[result.TrackId] = result
	}

	return m
}

type ItunesResponse struct {
	TrackId         int     `json:"trackId"`
	TrackName       string  `json:"trackName"`
	ArtistName      string  `json:"artistName"`
	TrackTimeMillis int     `json:"trackTimeMillis"`
	CollectionName  string  `json:"collectionName"`
	TrackPrice      float32 `json:"trackPrice"`
	ArtworkUrl      string  `json:"artworkUrl100"`
}
