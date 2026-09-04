package types

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// 4. CPF (11 dígitos, apenas números)
type CPF string

func (c CPF) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	clean := strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, string(c))

	if len(c) > 14 || len(clean) < 11 {
		return fmt.Errorf("erro de tamanho inválido")
	}

	return e.EncodeElement(clean, start)
}
