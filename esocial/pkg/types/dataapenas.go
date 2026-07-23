package types

import (
	"encoding/xml"
	"time"
)

// DataApenas (AAAA-MM-DD)
type DataApenas time.Time

func (d DataApenas) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	t := time.Time(d)
	if t.IsZero() {
		return nil
	}
	return e.EncodeElement(t.Format("2006-01-02"), start)
}

func (d *DataApenas) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := decoder.DecodeElement(&s, &start); err != nil {
		return err
	}
	parsed, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*d = DataApenas(parsed)
	return nil
}
