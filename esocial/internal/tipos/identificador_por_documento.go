package tipos

import (
	"encoding/xml"
	"errors"
)

type IdentificadorPorDocumento struct {
	Tipo XmlValueBase
}

func (ele IdentificadorPorDocumento) Validar() error {
	switch t := ele.Tipo.(type) {
	case *CNPJ:
		if err := t.Validar(); err != nil {
			return err
		}
	case *CPF:
		if err := t.Validar(); err != nil {
			return err
		}
	default:
		return errors.New("tipo inválido para IdeEmpregador")
	}
	return nil
}

func (ele IdentificadorPorDocumento) String() string {
	return ele.Tipo.String()
}

func (ele IdentificadorPorDocumento) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	err := ele.Validar()
	if err != nil {
		return err
	}

	start.Name.Local = "ideEmpregador"
	err = e.EncodeToken(start)
	if err != nil {
		return err
	}

	// Encode the Tipo field
	if ele.Tipo == nil {
		return errors.New("tipo não pode ser nulo")
	}

	if err := e.EncodeElement(ele.Tipo, xml.StartElement{Name: xml.Name{Local: "nrInsc"}}); err != nil {
		return err
	}

	var tpInsc int
	switch ele.Tipo.(type) {
	case *CNPJ:
		tpInsc = 1
	case *CPF:
		tpInsc = 2
	}

	if err := e.EncodeElement(tpInsc, xml.StartElement{Name: xml.Name{Local: "tpInsc"}}); err != nil {
		return err
	}

	err = e.EncodeToken(xml.EndElement{Name: start.Name})
	if err != nil {
		return err
	}

	return nil
}

func (ele *IdentificadorPorDocumento) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if err := d.DecodeElement(ele, &start); err != nil {
		return err
	}
	err := ele.Validar()
	if err != nil {
		return err
	}

	return nil
}
