package tipos

import (
	"encoding/xml"
	"fmt"
)

type NumRecibo struct {
	Ambiente         TipoAmbiente
	NumeroSequencial uint64
}

func (ele NumRecibo) Validar() error {
	if ele.NumeroSequencial == 0 || ele.NumeroSequencial > 9999999999999999999 {
		return fmt.Errorf("número sequencial inválido: %d", ele.NumeroSequencial)
	}
	return nil
}

func (ele NumRecibo) String() string {
	return fmt.Sprintf("1.%d.%019d", ele.Ambiente, ele.NumeroSequencial)
}

func (ele NumRecibo) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	err := ele.Validar()
	if err != nil {
		return err
	}

	value := ele.String()

	e.EncodeElement(value, start)
	return nil
}

func (ele NumRecibo) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var conteudo string
	if err := d.DecodeElement(&conteudo, &start); err != nil {
		return err
	}

	n, err := fmt.Sscanf(conteudo, "1.%d.%019d", &ele.Ambiente, &ele.NumeroSequencial)
	if err != nil || n != 2 {
		return fmt.Errorf("formato inválido para NumRecibo: %s", conteudo)
	}

	// Validar os dados após a leitura
	err = ele.Validar()
	if err != nil {
		return err
	}

	return nil
}
