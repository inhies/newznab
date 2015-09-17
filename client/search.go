package client

import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/inhies/newznab"
)

// Returned list of NZBs from the Indexer.
type SearchResults struct {
	// Total number of results found.
	Total int

	// How far in to all the found results we are.
	Offset int

	// NZBs matching the search query.
	NZBs []newznab.NZB

	// If the request returned an error this will be set.
	Error *newznab.Error
}

// A newznab SEARCH request. Only the Query field is required.
type SearchRequest struct {
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
}

// Perform a newznab SEARCH request on the specified Indexer.
func (i *Indexer) Search(req *SearchRequest) (*SearchResults, error) {
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
	var feed newznab.SearchResponse
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
func (i *Indexer) query(url *url.URL) ([]byte, *newznab.Error, error) {
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
	apiErr, httpErr := newznab.CheckForError(data)
	if apiErr != nil || httpErr != nil {
		return nil, apiErr, httpErr
	}

	//fmt.Printf("%s\n", data)
	return data, nil, nil
}
