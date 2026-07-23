package types

import (
	"encoding/xml"
	"time"
)

// DataAnoMes (AAAA-MM)
type DataAnoMes time.Time

func (d DataAnoMes) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	t := time.Time(d)
	if t.IsZero() {
		return nil
	}
	return e.EncodeElement(t.Format("2006-01"), start)
}

func (d *DataAnoMes) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := decoder.DecodeElement(&s, &start); err != nil {
		return err
	}
	parsed, err := time.Parse("2006-01", s)
	if err != nil {
		return err
	}
	*d = DataAnoMes(parsed)
	return nil
}
