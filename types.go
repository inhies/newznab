package newznab

import (
	"net/url"
	"time"

	"github.com/inhies/go-bytesize"
)

type TvNZB struct {
	NZB
	Attributes struct {
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
		Media            *Media
		Attributes
	} `xml:"attr"`
}
type MovieNZB struct {
	NZB
	Attributes struct {
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
		Media            *Media
		Attributes
	} `xml:"attr"`
}
type MusicNZB struct {
	NZB
	Attributes struct {
		Artist           string
		Album            string
		Publisher        string
		Tracks           string
		CoverURL         *url.URL
		BackdropCoverURL *url.URL
		Review           string
		Media            *Media
		Attributes
	} `xml:"attr"`
}
type BookNZB struct {
	NZB
	Attributes struct {
		Title       string
		Publisher   string
		PublishDate time.Time
		Author      string
		Pages       int
		CoverURL    *url.URL
		Review      string
		Attributes
	} `xml:"attr"`
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
	// [namespace][attribute name][attribute value]
	Unknown map[string]map[string]string
}

type Media struct {
	Video      string
	Audio      string
	Resolution string
	Framerate  string
	Language   string
	Subs       string
	Genre      string
}
