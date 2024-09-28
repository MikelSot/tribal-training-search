package model

type Entity string

const (
	Album       Entity = "album"
	MusicArtist Entity = "musicArtist"
	Song        Entity = "song"
	Artist      Entity = "artist"
)

type Search struct {
	Entity Entity
	Search string
}

func (s Search) IsSearchEmpty() bool { return s.Search == "" }

type SearchMap map[Entity]Search
