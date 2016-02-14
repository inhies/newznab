// Newznab provides easy access to compliant newznab usenet indexer APIs.
package newznab

import "net/http"

// TODO:REGISTER,DETAILS
// GETNFO,GET,CART-ADD,CART-DEL,COMMENTS,COMMENTS-ADD,USER
/*
func (i *Indexer) Caps() (*Capabilities, error) {
	return nil, nil
}
*/

// Indexer specific information.
type Indexer struct {
	// Name of the indexer, can be anything you like.
	Name string

	// Your API key to access the website.
	APIKey string

	// URL of the site. Include http:// or https://
	URL string

	// Enable use of this indexer.
	Enabled bool

	// Skip SSL certification verification. Some website use self signed
	// certificates and need this set to true.
	SkipSSLVerification bool

	client *http.Client
}

/*
// Searches the Indexer's television categories and allows you to specifiy
// the exact season and episode you want, as well as the TvRage ID number for
// the show.
func (i *Indexer) TvSearch(query *Query, season, episode string, TVRageID int) (*TvResults, error) {
	return nil, nil
}
func (i *Indexer) BookSearch(query *Query, title, author string) (*BookResults, error) {
	return nil, nil
}
func (i *Indexer) MusicSearch(query *Query, artist, album, label, track, genre string, year int) (*MusicResults, error) {
	return nil, nil
}
func (i *Indexer) MovieSearch(query *Query, genre string, IMDB int) (*MovieResults, error) {
	return nil, nil
}
*/

// NewQuery returns a new Query given a query string.
func NewQuery(query string) *Query {
	return &Query{Query: query}
}

// NewIndexer returns a new Indexer given a URL and API key.
func NewIndexer(url string, apikey string) *Indexer {
	return &Indexer{
		APIKey: apikey,
		URL:    url,
		client: &http.Client{
			Transport: http.DefaultTransport,
		},
	}
}

// REGISTER
// <register username="u675a9b6" password="dac20bde" apikey="43fc41e56e36db9d51fbfb2e717f1267"/>
