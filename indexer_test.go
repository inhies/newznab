package newznab

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func getTestServer(respCode int, body string) (*httptest.Server, *Indexer) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(respCode)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, body)
	}))

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	// Make a http.Client with the transport
	httpClient := &http.Client{Transport: transport}
	// Make an API client and inject
	indexer := &Indexer{URL: server.URL, client: httpClient}

	return server, indexer
}

func TestIndexer_Search_BadCredentials(t *testing.T) {
	server, indexer := getTestServer(200, errIncorrectCredentials)
	defer server.Close()
	results, err := indexer.Search(NewQuery("Public Domain Tv"))
	if err != nil {
		t.Fail()
	}
	if !reflect.DeepEqual(results.Error, ErrBadCreds) {
		t.Fail()
	}

}

func TestIndexer_Search_NoResults(t *testing.T) {
	server, indexer := getTestServer(200, searchResponseNoResults)
	defer server.Close()
	results, err := indexer.Search(NewQuery("Public Domain Tv"))
	if err != nil {
		t.Fail()
	}
	if results == nil {
		t.Fail()
	}
	if results.Total != 0 {
		t.Fail()
	}
}

func TestIndexer_Search(t *testing.T) {
	server, indexer := getTestServer(200, searchResponse)
	defer server.Close()
	results, err := indexer.Search(NewQuery("Public Domain Tv"))
	if err != nil {
		t.Fail()
	}
	if results.Total != 2344 {
		t.Log("Total results are not 2344:", results.Total)
		t.Fail()
	}
	if len(results.NZBs) != 1 {
		t.Log("NZBs returned != 1:", len(results.NZBs))
		t.Fail()
	}

}

var errIncorrectCredentials = `<error code="100" description="Incorrect user credentials"/>`

var searchResponseNoResults = `<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0" xmlns:newznab="http://www.newznab.com/DTD/2010/feeds/attributes/">
<channel>
    <newznab:response offset="0" total="0"/>
</channel>
</rss>`

var searchResponse = `<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0" xmlns:newznab="http://www.newznab.com/DTD/2010/feeds/attributes/">

<channel>
<title>example.com</title>
<description>example.com API results</description>
<!--
   More RSS content
 -->

<!-- offset is the current offset of the response
     total is the total number of items found by the query
 -->
<newznab:response offset="0" total="2344"/>

<item>
  <!-- Standard RSS 2.0 Data -->
  <title>A.Public.Domain.Tv.Show.S06E05</title>
  <guid isPermaLink="true">http://servername.com/rss/viewnzb/e9c515e02346086e3a477a5436d7bc8c</guid>
  <link>http://servername.com/rss/nzb/e9c515e02346086e3a477a5436d7bc8c&amp;i=1&amp;r=18cf9f0a736041465e3bd521d00a90b9</link>
  <comments>http://servername.com/rss/viewnzb/e9c515e02346086e3a477a5436d7bc8c#comments</comments>
  <pubDate>Sun, 06 Jun 2010 17:29:23 +0100</pubDate>
  <category>TV > XviD</category>
  <description>Some TV show</description>
  <enclosure url="http://servername.com/rss/nzb/e9c515e02346086e3a477a5436d7bc8c&amp;i=1&amp;r=18cf9f0a736041465e3bd521d00a90b9" length="154653309" type="application/x-nzb" />

  <!-- Additional attributes -->
  <newznab:attr name="category" value="2000"/>
  <newznab:attr name="category" value="2030"/>
  <newznab:attr name="size"     value="4294967295"/>
</item>
</channel>
</rss>`

/*
func ExampleIndexer_TvSearch() {
	i := NewIndexer("http://somesite.com", "abc123")
	q := NewQuery("public access television")
	r, err := i.TvSearch(q, "1", "1", 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	if r == nil {
		fmt.Println("nil returned")
		return
	}
}
*/
