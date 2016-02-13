package newznab

import (
	"encoding/xml"
	"time"
)

type Time struct {
	time.Time
}

func (t *Time) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(xml.CharData([]byte(t.UTC().Format(time.RFC1123Z))))
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
