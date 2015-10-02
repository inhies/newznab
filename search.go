package newznab

import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/inhies/go-bytesize"
	"github.com/kylelemons/godebug/pretty"
)

var Pretty = &pretty.Config{PrintStringers: true}

// Item represents a single NZB item in search results.
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

type attrs struct {
	XMLName xml.Name
	Name    string `xml:"name,attr"`
	Value   string `xml:"value,attr"`
}

func (a *Attributes) addUnknownAttr(attr *attrs) {
	if a.Unknown == nil {
		a.Unknown = make(map[string]map[string]string)
	}
	if a.Unknown[attr.XMLName.Space] == nil {
		a.Unknown[attr.XMLName.Space] = make(map[string]string)
	}
	a.Unknown[attr.XMLName.Space][attr.Name] = attr.Value
}
func (a *Attributes) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var raw attrs

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
				case "poster":
					a.Newznab.Poster = raw.Value
				case "group":
					a.Newznab.Group = raw.Value
				case "team":
					a.Newznab.Team = raw.Value
				case "grabs":
					i, err := strconv.Atoi(raw.Value)
					if err != nil {
						return err
					}
					a.Newznab.Grabs = i
				case "password":
				case "comments":
					i, err := strconv.Atoi(raw.Value)
					if err != nil {
						return err
					}
					a.Newznab.Comments = i
				case "usenetdate":
				case "info":
					u, err := url.Parse(raw.Value)
					if err != nil {
						return err
					}
					a.Newznab.Info = u
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
		a.addUnknownAttr(&raw)
	}
	//*t = Time{date}
	return nil

}

/*
func (i *Indexer) BookSearch(req *BookQuery) (*BookResults, error) {
	return nil, nil
}

type BookQuery struct {
	SearchQuery
	Title  string
	Author string
}
*/

type BookResults struct {
	// Total number of results found.
	Total int

	// How far in to all the found results we are.
	Offset int

	// NZBs matching the search query.
	NZBs []BookNZB

	// If the request returned an error this will be set.
	Error *Error
}

/*
func (i *Indexer) MusicSearch(req *MusicQuery) (*MusicResults, error) {
	return nil, nil
}

type MusicQuery struct {
	SearchQuery
	Album  string
	Artist string
	Label  string
	Track  string
	Year   int
	Genre  string
}
*/

type MusicResults struct {
	// Total number of results found.
	Total int

	// How far in to all the found results we are.
	Offset int

	// NZBs matching the search query.
	NZBs []MusicNZB

	// If the request returned an error this will be set.
	Error *Error
}

/*
func (i *Indexer) MovieSearch(req *MovieQuery) (*MovieResults, error) {
	return nil, nil
}

type MovieQuery struct {
	SearchQuery
	Genre  string
	IMDBID int
}
*/

type MovieResults struct {
	// Total number of results found.
	Total int

	// How far in to all the found results we are.
	Offset int

	// NZBs matching the search query.
	NZBs []MovieNZB

	// If the request returned an error this will be set.
	Error *Error
}

/*
func (i *Indexer) TvSearch(req *TvQuery) (*TvResults, error) {
	return nil, nil
}

type TvQuery struct {
	SearchQuery
	TVRageID int
	Season   string
	Episode  string
}

*/
type TvResults struct {
	// Total number of results found.
	Total int

	// How far in to all the found results we are.
	Offset int

	// NZBs matching the search query.
	NZBs []TvNZB

	// If the request returned an error this will be set.
	Error *Error
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
type SearchQuery struct {
	// The search query.
	Query string

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

	indexer *Indexer
}

// Perform a newznab SEARCH request on the specified Indexer.
func (i *Indexer) search(req *SearchQuery) (*SearchResults, error) {
	// URL encode the search string.
	v := url.Values{}
	v.Add("t", "search")
	v.Set("apikey", i.APIKey)
	v.Add("o", "xml") // XML by default please, becase JSON doesn't exist!
	v.Add("q", req.Query)

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
			str = append(str, string(int(value)))
		}
		v.Add("cat", strings.Join(str, ","))
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

	finalURL.RawQuery = v.Encode()

	data, apiErr, httpErr := i.query(finalURL)
	if httpErr != nil {
		return nil, err
	}

	r := new(SearchResults)
	if apiErr != nil {
		r.Error = apiErr
		return nil, nil
	}

	// No error found, unmarsal the returned RSS feed
	var feed SearchResponse
	err = xml.Unmarshal(data, &feed)
	if err != nil {
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
	return r, nil
}

// figure out a way to do errors better
func (i *Indexer) query(url *url.URL) ([]byte, *Error, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: i.SkipSSLVerification},
	}

	client := &http.Client{Transport: tr}
	fmt.Println("GET", url.String(), "\n")
	res, err := client.Get(url.String())
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
	apiErr, httpErr := CheckForError(data)
	if apiErr != nil || httpErr != nil {
		return nil, apiErr, httpErr
	}

	//fmt.Printf("%s\n", data)
	return data, nil, nil
}
