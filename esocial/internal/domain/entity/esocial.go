package entity

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"reflect"
	"strings"
)

const (
	erro_xml_tag = "o campo XMLName está vazio"
)

type ESocial struct {
	Versao string
	Evento any
}

func (a ESocial) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if tipo, ok := reflect.TypeOf(a.Evento).FieldByName("XMLName"); !ok {
		return fmt.Errorf("não foi possível encontrar o campo XMLName no evento")
	} else {
		xml_name := tipo.Tag.Get("xml")
		if xml_name == "" {
			return fmt.Errorf(erro_xml_tag)
		}
		xml_name = strings.Split(xml_name, ",")[0]
		if xml_name == "" {
			return fmt.Errorf(erro_xml_tag)
		}
		name_space := fmt.Sprintf("http://www.esocial.gov.br/schema/evt/%v/%s", xml_name, a.Versao)
		start.Name = xml.Name{Local: "eSocial", Space: name_space}
		e.EncodeToken(start)
		e.Encode(a.Evento)
		e.EncodeToken(start.End())
	}

	return nil
}

func (a ESocial) ToXML() (string, error) {
	data, err := xml.Marshal(a)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	buf.WriteString(xml.Header)
	buf.WriteString(string(data))
	return buf.String(), nil
}
