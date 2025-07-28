package tipos

import (
	"encoding/xml"
	"fmt"
)

// Número de recibo do evento.
// Abaixo é descrita a regra de
// formação deste código:

// A.B.NNNNNNNN....N

// A = Agente de processamento:
// Serpro=1
// B = Ambiente de recepção:
// 1=Produção;
// 2=Pré-produção - dados reais;
// 3=Pré-produção - dados
// fictícios;
// 6=Homologação; 7=Validação;
// 8=Testes;
// 9=Desenvolvimento;
// N = Número sequencial (19
// posições)

type NumRecibo struct {
	Ambiente         TipoAmbiente
	NumeroSequencial uint64
}

func (ele NumRecibo) Validar() error {
	if ele.NumeroSequencial == 0 || ele.NumeroSequencial > 9999999999999999999 {
		return fmt.Errorf("Número sequencial inválido: %d", ele.NumeroSequencial)
	}
	return nil
}

func (ele NumRecibo) String() string {
	return fmt.Sprintf("1.%d.%d", ele.Ambiente, ele.NumeroSequencial)
}

func (ele NumRecibo) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	err := ele.Validar()
	if err != nil {
		return err
	}

	//TODO: Implementar

	return nil
}

func (ele NumRecibo) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var conteudo string
	if err := d.DecodeElement(&conteudo, &start); err != nil {
		return err
	}

	//TODO: Implementar

	err := ele.Validar()
	if err != nil {
		return err
	}

	return nil
}
