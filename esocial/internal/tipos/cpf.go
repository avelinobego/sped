package tipos

import (
	"encoding/xml"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type CPF string

func (ele CPF) Tipo() int {
	return 2
}

func (ele CPF) Validar() error {
	cpf := string(ele)
	re := regexp.MustCompile(`\D`)
	cpf = re.ReplaceAllString(cpf, "")

	if len(cpf) != 11 {
		return errors.New("CPF deve conter 11 dígitos")
	}

	for i := 0; i < 10; i++ {
		if cpf == strings.Repeat(strconv.Itoa(i), 11) {
			return errors.New("CPF inválido: todos os dígitos são iguais")
		}
	}

	encontrado := 0
	for t := 9; t < 11; t++ {
		sum := 0
		for i := 0; i < t; i++ {
			num, _ := strconv.Atoi(string(cpf[i]))
			sum += num * (t + 1 - i)
		}
		dv := (sum * 10) % 11
		if dv == 10 {
			dv = 0
		}

		encontrado = int(cpf[t] - '0')

		if dv != encontrado {
			return fmt.Errorf("CPF inválido: dígito verificador %d incorreto. esperado %d", encontrado, dv)
		}
	}

	return nil
}

func (ele CPF) String() string {
	return string(ele)
}

func (ele CPF) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	err := ele.Validar()
	if err != nil {
		return err
	}
	err = e.EncodeElement(string(ele), start)
	if err != nil {
		return err
	}

	return nil
}

func (ele *CPF) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var conteudo string
	if err := d.DecodeElement(&conteudo, &start); err != nil {
		return err
	}

	*ele = CPF(conteudo)

	err := ele.Validar()
	if err != nil {
		return err
	}

	return nil
}
