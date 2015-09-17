package newznab

import (
	"encoding/xml"
	"time"
)

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
		Value string `xml:"url,chardata"`
	} `xml:"source,omitempty"`

	// A custom time.Time wrapper to enable direct marshalling and
	// unmarshalling. Still need to write the marshaller...
	Date Time `xml:"pubDate",omitempty`

	Enclosure struct {
		URL    string `xml:"url,attr"`
		Length string `xml:"length,attr"`
		Type   string `xml:"type,attr"`
	} `xml:"enclosure",omitempty`

	// Need to write a marshaller
	Attributes []struct {
		XMLName xml.Name
		Name    string `xml:"name,attr"`
		Value   string `xml:"value,attr"`
	} `xml:"attr"`
}

type Time struct {
	time.Time
}

func (t *Time) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(xml.CharData([]byte(t.UTC().Format(time.RFC822))))
	e.EncodeToken(xml.EndElement{start.Name})
	return nil
}

func (t *Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var raw string

	err := d.DecodeElement(&raw, &start)
	if err != nil {
		return err
	}
	date, err := time.Parse(time.RFC1123Z, raw)

	if err != nil {
		return err
	}

	*t = Time{date}
	return nil

}
