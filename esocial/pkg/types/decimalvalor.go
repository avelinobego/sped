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
	format := fmt.Sprintf("%%.%df", v.Casas)
	formatted := fmt.Sprintf(format, v.Value)
	// O eSocial exige ponto como separador decimal
	formatted = strings.ReplaceAll(formatted, ",", ".")
	return e.EncodeElement(formatted, start)
}
