package types_test

import (
	"encoding/xml"
	"sped/esocial/pkg/types"
)

type WrapperAll struct {
	XMLName    xml.Name           `xml:"wrapper"`
	DataAnoMes types.DataAnoMes   `xml:"dataAnoMes,omitempty"`
	DataApenas types.DataApenas   `xml:"dataApenas,omitempty"`
	Decimal    types.DecimalValor `xml:"decimal,omitempty"`
	CPF        types.CPF          `xml:"cpf,omitempty"`
	SimNao     types.SimNao       `xml:"simNao,omitempty"`
}
