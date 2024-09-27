package model

import "encoding/xml"

type ChartLyricsResult struct {
	SearchLyricResult []ChartLyricsResponse `xml:"SearchLyricResult"`
}

func (c ChartLyricsResult) IsEmpty() bool { return len(c.SearchLyricResult) == 0 }

type ChartLyricsResponse struct {
	XMLName       xml.Name `xml:"SearchLyricResult"`
	TrackChecksum string   `xml:"TrackChecksum"`
	TrackId       int      `xml:"TrackId"`
	LyricChecksum string   `xml:"LyricChecksum"`
	LyricId       int      `xml:"LyricId"`
	SongUrl       string   `xml:"SongUrl"`
	ArtistUrl     string   `xml:"ArtistUrl"`
	Artist        string   `xml:"Artist"`
	Song          string   `xml:"Song"`
	SongRank      int      `xml:"SongRank"`
}
