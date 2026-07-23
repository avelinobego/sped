package types

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// CNPJ Alfanumérico (14 caracteres) com validação de tamanho e dígitos verificadores (Módulo 11)
type CNPJ string

func (c CNPJ) validar() (string, error) {
	clean := strings.Map(func(r rune) rune {
		if (r >= '0' && r <= '9') || (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
			if r >= 'a' && r <= 'z' {
				return r - 32
			}
			return r
		}
		return -1
	}, string(c))

	if len(clean) != 14 {
		return "", fmt.Errorf("tamanho inválido para CNPJ: %d (esperado 14)", len(clean))
	}

	// Cálculo do primeiro dígito verificador
	pesos1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	soma1 := 0
	for i := range 12 {
		val := int(clean[i]) - 48
		soma1 += val * pesos1[i]
	}
	resto1 := soma1 % 11
	dv1 := 11 - resto1
	if dv1 >= 10 {
		dv1 = 0
	}

	// Cálculo do segundo dígito verificador
	pesos2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	soma2 := 0
	for i := range 12 {
		val := int(clean[i]) - 48
		soma2 += val * pesos2[i]
	}
	soma2 += dv1 * pesos2[12]
	resto2 := soma2 % 11
	dv2 := 11 - resto2
	if dv2 >= 10 {
		dv2 = 0
	}

	// Validação dos dígitos informados versus calculados
	dvInformado := int(clean[12]-48)*10 + int(clean[13]-48)
	dvCalculado := dv1*10 + dv2

	if dvInformado != dvCalculado {
		return "", fmt.Errorf("dígitos verificadores inválidos para o CNPJ %s", clean)
	}

	return clean, nil
}

func (c CNPJ) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	clean, err := c.validar()
	if err != nil {
		return err
	}
	return e.EncodeElement(clean, start)
}

func (c *CNPJ) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := decoder.DecodeElement(&s, &start); err != nil {
		return err
	}

	temp := CNPJ(s)
	clean, err := temp.validar()
	if err != nil {
		return err
	}

	*c = CNPJ(clean)
	return nil
}
