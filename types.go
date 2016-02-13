package newznab

import (
	"net/url"
	"time"

	"github.com/inhies/go-bytesize"
)

/*
type SearchType int

const (
	StandardSearch SearchType = iota + 1
	TvSearch
	MovieSearch
	MusicSearch
	BookSearch
)
*/

type TvAttrs struct {
	Season           int
	Episode          int
	RageID           int
	TVDBID           int
	Rating           int
	TVTitle          string
	TVAirDate        time.Time
	CoverURL         *url.URL
	BackdropCoverURL *url.URL
	Review           string
	Media            *MediaAttrs
}

type MovieAttrs struct {
	IMDBID           int // IMDB ID
	Score            int
	Title            string
	Tagline          string
	Plot             string
	Year             int
	Director         string
	Actors           string
	CoverURL         *url.URL
	BackdropCoverURL *url.URL
	Review           string
	Media            *MediaAttrs
}
type MusicAttrs struct {
	Artist           string
	Album            string
	Publisher        string
	Tracks           string
	CoverURL         *url.URL
	BackdropCoverURL *url.URL
	Review           string
	Media            *MediaAttrs
}
type BookAttrs struct {
	Title       string
	Publisher   string
	PublishDate time.Time
	Author      string
	Pages       int
	CoverURL    *url.URL
	Review      string
}

type MediaAttrs struct {
	Video      string
	Audio      string
	Resolution string
	Framerate  string
	Language   string
	Subs       string
	Genre      string
}

type Attributes struct {
	Size         bytesize.ByteSize
	Categories   []int
	GUID         string
	Files        int
	Poster       string
	Group        string
	Team         string
	Grabs        int
	Passworded   bool
	InnerArchive bool
	Comments     int
	UsenetDate   time.Time
	Info         *url.URL
	Year         int

	// Other attributes added by third parties, accessed via:
	// [namespace][attribute name]
	// currently broken, not sure if I even want to support this
	//Unknown map[string]*Attribute

	// Tv-Search specific attributes.
	Tv *TvAttrs `xml:",omitempty"`

	// Movie-Search specific attributes.
	Movie *MovieAttrs `xml:",omitempty"`

	// Music-Search specific attributes.
	Music *MusicAttrs `xml:",omitempty"`

	// Attributes relating to the specific media formats and files.
	Media *MediaAttrs `xml:",omitempty"`

	// Book-Search specific attributes.
	Book *BookAttrs `xml:",omitempty"`
}
