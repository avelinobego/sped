package tipos

import "encoding/xml"

type Cpf int

func (ele Cpf) Validar() error {
	return nil
}

func (ele Cpf) String() string {
	return ""
}

func (ele Cpf) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	err := ele.Validar()
	if err != nil {
		return err
	}
	//TODO: Implementar
	return nil
}

func (ele *Cpf) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	err := ele.Validar()
	if err != nil {
		return err
	}
	//TODO: Implementar
	return nil
}
