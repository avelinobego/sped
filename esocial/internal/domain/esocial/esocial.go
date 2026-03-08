package esocial

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

func NewESocial(versao string, evento any) *eSocial {
	return &eSocial{
		versao: versao,
		evento: evento,
	}
}

type eSocial struct {
	versao string
	evento any
	buffer *bytes.Buffer
}

func (a *eSocial) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if tipo, ok := reflect.TypeOf(a.evento).FieldByName("XMLName"); !ok {
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
		name_space := fmt.Sprintf("http://www.esocial.gov.br/schema/evt/%v/%s", xml_name, a.versao)
		start.Name = xml.Name{Local: "eSocial", Space: name_space}
		e.EncodeToken(start)
		e.Encode(a.evento)
		e.EncodeToken(start.End())
	}

	return nil
}

func (a *eSocial) ToXML() (string, error) {
	data, err := xml.Marshal(a)
	if err != nil {
		return "", err
	}

	if a.buffer == nil {
		a.buffer = &bytes.Buffer{}
		a.buffer.WriteString(xml.Header)
		a.buffer.Truncate(a.buffer.Len() - 1)
		a.buffer.Write(data)
	}
	return a.buffer.String(), nil
}
