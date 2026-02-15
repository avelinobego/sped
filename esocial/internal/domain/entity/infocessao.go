package entity

type Infocessao struct {
	Cnpjcess  string      `xml:"cnpjCess"`
	Infonus   int64       `xml:"infOnus"`
	Inicessao []Inicessao `xml:"iniCessao,omitempty"`
	Fimcessao []Fimcessao `xml:"fimCessao,omitempty"`
}
