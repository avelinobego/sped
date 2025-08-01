package tipos

import "encoding/xml"

type XmlValueBase interface {
	Validar() error
	String() string
	MarshalXML(e *xml.Encoder, start xml.StartElement) error
	UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
}
