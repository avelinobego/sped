package types

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// 3. DecimalValor (Valores monetários ou numéricos com casas decimais fixas)
type DecimalValor struct {
	Value float64
	Casas int
}

func (v DecimalValor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	if v.Value == 0 && v.Casas == 0 {
		return e.EncodeElement(nil, start)
	}

	if v.Casas < 0 {
		return fmt.Errorf("casas decimais não podem ser negativas")
	}

	format := fmt.Sprintf("%%.%df", v.Casas)
	formatted := fmt.Sprintf(format, v.Value)
	// O eSocial exige ponto como separador decimal
	formatted = strings.ReplaceAll(formatted, ",", ".")
	return e.EncodeElement(formatted, start)
}
