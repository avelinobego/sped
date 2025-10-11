package tipos

import (
	"encoding/xml"
	"time"
)

type IdePeriodo time.Time

func (ip IdePeriodo) String() string {
	return time.Time(ip).Format("2006-01")
}

func (ip IdePeriodo) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(time.Time(ip).Format("2006-01"), start)
}

func (ip *IdePeriodo) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var conteudo string
	if err := d.DecodeElement(&conteudo, &start); err != nil {
		return err
	}
	value, err := time.Parse("2006-01", conteudo)
	if err != nil {
		return err
	}
	*ip = IdePeriodo(value)
	return nil
}
