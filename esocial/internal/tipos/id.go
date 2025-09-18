package tipos

import (
	"encoding/xml"
	"fmt"
	"time"
)

type Id string

func NewId[T CpfCnpjConstraint](cpfCnpj T, dataHora time.Time, sequencia int) (result Id, err error) {
	err = cpfCnpj.Validar()
	if err != nil {
		return
	}

	result = Id(fmt.Sprintf("ID%d%014s%14s%05d",
		cpfCnpj.Tipo(),
		string(cpfCnpj),
		dataHora.Format("20060102150405"),
		sequencia))

	return
}

func (ele Id) Validar() error {
	return nil
}

func (ele Id) String() string {
	return string(ele)
}

func (ele Id) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	err := ele.Validar()
	if err != nil {
		return err
	}

	return e.EncodeElement(ele.String(), start)
}

func (ele *Id) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var conteudo string
	if err := d.DecodeElement(&conteudo, &start); err != nil {
		return err
	}

	err := ele.Validar()
	if err != nil {
		return err
	}

	return nil
}
