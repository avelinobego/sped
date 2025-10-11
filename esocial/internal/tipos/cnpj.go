package tipos

import (
	"encoding/xml"
	"errors"
)

type CNPJ string

func (ele CNPJ) Tipo() int {
	return 1
}

func (ele CNPJ) Validar() error {
	cnpj := string(ele)
	if len(cnpj) != 14 {
		return errors.New("CNPJ deve ter 14 dígitos")
	}

	// Check if all digits are the same
	allEqual := true
	for i := 1; i < len(cnpj); i++ {
		if cnpj[i] != cnpj[0] {
			allEqual = false
			break
		}
	}
	if allEqual {
		return errors.New("CNPJ inválido")
	}

	// First digit validation
	sum := 0
	weight := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	for i := 0; i < 12; i++ {
		digit := int(cnpj[i] - '0')
		sum += digit * weight[i]
	}
	remainder := sum % 11
	calculatedDigit := 0
	if remainder >= 2 {
		calculatedDigit = 11 - remainder
	}
	if int(cnpj[12]-'0') != calculatedDigit {
		return errors.New("CNPJ inválido")
	}

	// Second digit validation
	sum = 0
	weight = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	for i := 0; i < 13; i++ {
		digit := int(cnpj[i] - '0')
		sum += digit * weight[i]
	}
	remainder = sum % 11
	calculatedDigit = 0
	if remainder >= 2 {
		calculatedDigit = 11 - remainder
	}
	if int(cnpj[13]-'0') != calculatedDigit {
		return errors.New("CNPJ inválido")
	}

	return nil
}

func (ele CNPJ) String() string {
	return string(ele)
}

func (ele CNPJ) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	err := ele.Validar()
	if err != nil {
		return err
	}

	return e.EncodeElement(ele.String(), start)
}

func (ele *CNPJ) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var conteudo string
	if err := d.DecodeElement(&conteudo, &start); err != nil {
		return err
	}

	*ele = CNPJ(conteudo)

	err := ele.Validar()
	if err != nil {
		return err
	}

	return nil
}
