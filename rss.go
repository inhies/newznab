package newznab

// Search requests are responded to with an RSS feed.
type SearchResponse struct {
	// RSS version of the response.
	Version string `xml:"version,attr"`
	Channel struct {
		Title string `xml:"title"`
		Link  struct {
			Href string `xml:"href,attr"`
			Rel  string `xml:"rel,attr"`
			Type string `xml:"type,attr"`
		} `xml:"http://www.w3.org/2005/Atom link"`
		Description string `xml:"description"`
		Language    string `xml:"language,omitempty"`
		Webmaster   string `xml:"webmaster,omitempty"`
		Category    string `xml:"category,omitempty"`
		Image       struct {
			URL         string `xml:"url"`
			Title       string `xml:"title"`
			Link        string `xml:"link"`
			Description string `xml:"description,omitempty"`
			Width       int    `xml:"width,omitempty"`
			Height      int    `xml:"height,omitempty"`
		} `xml:"image"`

		// Newznab specific attribute that lists the total number of
		// results found and the zero based offset of the results that
		// were returned.
		Response struct {
			Offset int `xml:"offset,attr"`
			Total  int `xml:"total,attr"`
		} `xml:"http://www.newznab.com/DTD/2010/feeds/attributes/ response"`

		// All NZBs that match the search query, up to the response limit.
		NZBs []NZB `xml:"item"`
		/*
			// Extra RSS fields that we dont need to worry about yet.
			Copyright   string
			Editor      string `xml:"managingEditor"`
			LastBuilt   Time
			Generator   string
			Docs        string
			Cloud       string
			TTL         int
			SkipHours   struct {
				Hours []int `xml:"hour"`
			}
			SkipDays struct {
				Days []string `xml:"day"`
			}
		*/
	} `xml:"channel"`
}
