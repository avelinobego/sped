package types

import (
	"encoding/xml"
	"fmt"
)

// 5. SimNao (Indicadores binários: 1 - Sim, 2 - Não)
type SimNao int

const (
	Sim SimNao = 1
	Nao SimNao = 2
)

func (s SimNao) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if s != 1 && s != 2 {
		return fmt.Errorf("valor inválido para SimNao: %d", s)
	}
	return e.EncodeElement(int(s), start)
}
