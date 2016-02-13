package newznab

import (
	"crypto/tls"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/inhies/go-bytesize"
	"github.com/kylelemons/godebug/pretty"
)

var Pretty = &pretty.Config{PrintStringers: true}

// NZB represents a single NZB item in search results.
type NZB struct {
	Title    string `xml:"title",omitempty`
	Link     string `xml:"link",omitempty`
	Category struct {
		Domain string `xml"domain,attr"`
		Value  string `xml:",chardata"`
	} `xml:"category",omitempty`

	GUID struct {
		GUID        string `xml:",chardata"`
		IsPermaLink bool   `xml:"isPermaLink,attr"`
	} `xml:"guid,omitempty"`

	Comments    string `xml:"comments"`
	Description string `xml:"description"`
	Author      string `xml:"author,omitempty"`

	// The original source for the item.
	Source struct {
		URL   string `xml:"url,attr"`
		Value string `xml:",chardata"`
	} `xml:"source,omitempty"`

	// A custom time.Time wrapper to enable direct marshalling and
	// unmarshalling. Still need to write the marshaller...
	PubDate Time `xml:"pubDate",omitempty`

	Enclosure struct {
		URL    string `xml:"url,attr"`
		Length string `xml:"length,attr"`
		Type   string `xml:"type,attr"`
	} `xml:"enclosure",omitempty`

	Attributes Attributes `xml:"attr"`
}

type Attribute struct {
	XMLName xml.Name
	Name    string `xml:"name,attr"`
	Value   string `xml:"value,attr"`
}

/*
func (a *Attributes) addUnknownAttr(attr *Attribute) {
	if a.Unknown == nil {
		a.Unknown = make(map[string]*Attribute)
	}
	if a.Unknown[attr.XMLName.Space] == nil {
		a.Unknown[attr.XMLName.Space] = &Attribute{}
	}
	a.Unknown[attr.XMLName.Space] = &Attribute{
		Name:  attr.Name,
		Value: attr.Value,
	}
}
*/
func (a *Attributes) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var raw Attribute

	err := d.DecodeElement(&raw, &start)
	if err != nil {
		return err
	}

	if raw.XMLName.Space == "http://www.newznab.com/DTD/2010/feeds/attributes/" &&
		raw.Name != "" && raw.Value != "" {
		switch strings.ToLower(raw.Name) {
		case "size":
			b, err := bytesize.Parse(raw.Value + "B")
			if err != nil {
				return err
			}
			a.Size = b
			//a.Newznab.Info = &url.URL{}
			/*
				case "category":
					i, err := strconv.Atoi(raw.Value)
					if err != nil {
						return err
					}
					a.Newznab.Categories = append(a.Newznab.Categories, i)
				case "guid":
					a.Newznab.GUID = raw.Value
				case "files":
					i, err := strconv.Atoi(raw.Value)
					if err != nil {
						return err
					}
					a.Newznab.Files = i
			*/
		case "poster":
			a.Poster = raw.Value
			/*
				case "group":
					a.Newznab.Group = raw.Value
				case "team":
					a.Newznab.Team = raw.Value
			*/
		case "grabs":
			i, err := strconv.Atoi(raw.Value)
			if err != nil {
				return err
			}
			a.Grabs = i
		case "password":
		case "comments":
			i, err := strconv.Atoi(raw.Value)
			if err != nil {
				return err
			}
			a.Comments = i
		case "usenetdate":
		case "info":
			u, err := url.Parse(raw.Value)
			if err != nil {
				return err
			}
			a.Info = u
			/*
				case "year":

				// Start of TV specific items
				case "season":
					i, err := strconv.Atoi(raw.Value)
					if err != nil {
						return err
					}
					n.Season = i
				case "episode":
				case "tvdbid":
					i, err := strconv.Atoi(raw.Value)
					if err != nil {
						return err
					}
					a.Newznab.TVDBID = i
				case "rageid":
					i, err := strconv.Atoi(raw.Value)
					if err != nil {
						return err
					}
					a.Newznab.RageID = i
				case "tvtitle":
					a.Newznab.TVTitle = raw.Value
				case "tvairdate":
				case "video":
					a.Newznab.Video = raw.Value
				case "audio":
					a.Newznab.Audio = raw.Value
				case "resolution":
				case "framerate":
				case "language":
					a.Newznab.Language = raw.Value
				case "subs":
					a.Newznab.Subs = raw.Value
				case "imdb":
					i, err := strconv.Atoi(raw.Value)
					if err != nil {
						return err
					}
					a.Newznab.IMDB = i
				case "imdbscore":
					i, err := strconv.Atoi(raw.Value)
					if err != nil {
						return err
					}
					a.Newznab.IMDBScore = i
				case "imdbtitle":
					a.Newznab.IMDBTitle = raw.Value
				case "imdbtagline":
					a.Newznab.IMDBTagline = raw.Value
				case "imdbplot":
					a.Newznab.IMDBPlot = raw.Value
				case "imdbyear":
				case "imdbdirector":
					a.Newznab.IMDBDirector = raw.Value
				case "imdbactors":
					a.Newznab.IMDBActors = raw.Value
				case "genre":
					a.Newznab.Genre = raw.Value
				case "artist":
					a.Newznab.Artist = raw.Value
				case "album":
					a.Newznab.Album = raw.Value
				case "publisher":
					a.Newznab.Publisher = raw.Value
				case "tracks":
					a.Newznab.Tracks = raw.Value
				case "coverurl":
					u, err := url.Parse(raw.Value)
					if err != nil {
						return err
					}
					a.Newznab.CoverURL = u
				case "backdropcoverurl":
					u, err := url.Parse(raw.Value)
					if err != nil {
						return err
					}
					a.Newznab.BackdropCoverURL = u
				case "review":
					a.Newznab.Review = raw.Value
				case "booktitle":
					a.Newznab.BookTitle = raw.Value
				case "publishdate":
				case "author":
					a.Newznab.Author = raw.Value
				case "pages":
					i, err := strconv.Atoi(raw.Value)
					if err != nil {
						return err
					}
					a.Newznab.Pages = i
			*/
		default:
		}
	} else {
		//a.addUnknownAttr(&raw)
	}
	//*t = Time{date}
	return nil

}

// Returned list of NZBs from the Indexer.
type SearchResults struct {
	// Total number of results found.
	Total int

	// How far in to all the found results we are.
	Offset int

	// NZBs matching the search query.
	NZBs []NZB

	// If the request returned an error this will be set.
	Error *Error
}

// A newznab SEARCH request. Only the Query field is required.
type Query struct {
	// The search query.
	Query string

	// search, tvsearch, etc
	SearchType string

	// Limit search to these newsgroups.
	Groups []string

	// Upper limit for the number of items to be returned.
	Limit int

	// Limit the search to these newznab categories.
	Categories []int

	// List of requested extended attributes.
	Attributes []string

	// Return all extended attributes and ignore the Attributes field.
	Extended bool

	// Delete the item from a users cart on download.
	Delete bool

	// Only return results which were posted to usenet in the last x days.
	MaxAge int

	// The 0 based query offset defining which part of the response we want.
	Offset int

	// After a Query has been executed, Feed will be set to the returned
	// RSS feed data.
	Feed *RSS

	Params map[string]string
}

// Default search query parameters.
var DefaultQuery = &Query{
	SearchType: "search",
	Limit:      100,
	Extended:   true,
}

func NewQueryFromURL(URL string) (*Query, error) {
	q := DefaultQuery
	parsed, err := url.Parse(URL)
	if err != nil {
		return nil, err
	}

	values := parsed.Query()
	for k, v := range values {
		switch k {
		case "t":
			q.SearchType = strings.Join(v, ",")
		case "q":
			q.Query = strings.Join(v, ",")
		case "group":
			q.Groups = v
		case "limit":
			q.Limit, err = strconv.Atoi(v[len(v)-1])
			if err != nil {
				return nil, err
			}
		case "cat":
			//q.Query = strings.Join(v, ",")
		case "o":
			//q.Query = strings.Join(v, ",")
		case "attrs":
			q.Attributes = v
		case "extended":
			for _, x := range v {
				if x == "1" {
					q.Extended = true
				}
			}
		case "del":
			for _, x := range v {
				if x == "1" {
					q.Delete = true
				}
			}
		case "maxage":
			//q.Query = strings.Join(v, ",")
		case "offset":
			q.Offset, err = strconv.Atoi(v[len(v)-1])
			if err != nil {
				return nil, err
			}
		case "apikey":
			continue
		default:
			if q.Params == nil {
				q.Params = make(map[string]string)
			}
			q.Params[k] = strings.Join(v, ",")
		}
	}
	return q, nil
}

// GetQueryURL returns the complete API request URL that the Indexer would
// request when performing a search.
func (i *Indexer) GetQueryURL(req *Query) (*url.URL, error) {
	// URL encode the search string.
	v := url.Values{}
	v.Add("t", req.SearchType)
	v.Set("apikey", i.APIKey)
	v.Add("o", "xml")
	if req.Query != "" {
		v.Add("q", req.Query)
	}

	if len(req.Groups) > 0 {
		v.Add("group", strings.Join(req.Groups, ","))
	}

	if req.Limit > 0 {
		v.Add("limit", strconv.Itoa(req.Limit))
	}

	// Convert category slice to a string with a comma between entries.
	if len(req.Categories) > 0 {
		var str []string
		for _, value := range req.Categories {
			str = append(str, strconv.Itoa(value))
		}
		v.Set("cat", strings.Join(str, ","))
	}

	if req.Extended {
		v.Add("extended", "1")
	} else if len(req.Attributes) > 0 {
		var str []string
		for _, value := range req.Attributes {
			str = append(str, string(value))
		}
		v.Add("attrs", strings.Join(str, ","))
	}

	if req.Delete {
		v.Add("del", "1")
	}

	if req.MaxAge > 0 {
		v.Add("maxage", strconv.Itoa(req.MaxAge))
	}

	if req.Offset > 0 {
		v.Add("offset", strconv.Itoa(req.Offset))
	}

	finalURL, err := url.Parse(i.URL + "/api")
	if err != nil {
		return nil, err
	}

	for k, p := range req.Params {
		v.Add(k, p)
	}
	finalURL.RawQuery = v.Encode()
	return finalURL, nil
}

// Execute a search in all categories. Most times this will be the best option
// if you don't need to include extra filtering of the results. For example,
// when searching for a Tv show but you don't care about the season or episode.
func (i *Indexer) Search(req *Query) (*SearchResults, error) {
	finalURL, err := i.GetQueryURL(req)
	if err != nil {
		return nil, err
	}

	data, apiErr, httpErr := i.query(finalURL)
	if httpErr != nil {
		return nil, err
	}

	r := new(SearchResults)
	if apiErr != nil {
		r.Error = apiErr
		return r, nil
	}

	// No error found, unmarsal the returned RSS feed
	var feed = new(RSS)
	//feed.Channel = SearchResults{}

	err = xml.Unmarshal(data, &feed)
	if err != nil {
		pretty.Print(string(data))
		return nil, err
	}
	/*
		pretty.Print(feed)
		output, err := xml.MarshalIndent(feed, "  ", "    ")
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		fmt.Printf("%s\n", output)
	*/
	r.NZBs = feed.Channel.NZBs
	r.Total = feed.Channel.Response.Total
	r.Offset = feed.Channel.Response.Offset

	if len(r.NZBs) == 0 || r.Total == 0 {
		//return nil, nil
	}
	req.Feed = feed
	return r, nil
}

// figure out a way to do errors better
func (i *Indexer) query(url *url.URL) ([]byte, *Error, error) {
	if i.client == nil {
		i.client = &http.Client{
			Transport: http.DefaultTransport,
		}
	}

	if i.client.Transport == nil {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: i.SkipSSLVerification},
		}

		//client := &http.Client{Transport: tr}
		i.client.Transport = tr
	}
	res, err := i.client.Get(url.String())
	if err != nil {
		return nil, nil, err
	}

	var data []byte

	data, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, nil, err
	}

	// Check for an error condition.
	apiErr, httpErr := checkForError(data)
	if apiErr != nil || httpErr != nil {
		return nil, apiErr, httpErr
	}

	return data, nil, nil
}
