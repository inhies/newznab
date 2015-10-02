package newznab

// TODO:REGISTER,DETAILS
// GETNFO,GET,CART-ADD,CART-DEL,COMMENTS,COMMENTS-ADD,USER
func (i *Indexer) Caps() (*Capabilities, error) {
	return nil, nil
}

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
}

func (s *SearchQuery) BookSearch(title, author string) (*BookResults, error) { return nil, nil }
func (s *SearchQuery) MusicSearch(artist, album, label, track, genre string, year int) (*MusicResults, error) {
	return nil, nil
}
func (s *SearchQuery) MovieSearch(genre string, IMDB int) (*MovieResults, error) { return nil, nil }
func (s *SearchQuery) TvSearch(season, episode string, TVRageID int) (*TvResults, error) {
	return nil, nil
}
func (s *SearchQuery) Search() (*SearchResults, error) { return nil, nil }
func (i *Indexer) NewQuery(query string) *SearchQuery  { return nil }
func New(url string, apikey string) (*Indexer, error) {
	i := &Indexer{
		APIKey: apikey,
		URL:    url,
	}
	return i, nil
}

// REGISTER
// <register username="u675a9b6" password="dac20bde" apikey="43fc41e56e36db9d51fbfb2e717f1267"/>
