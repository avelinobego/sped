package entity

type Cargofuncao struct {
	Nmcargo   []string `xml:"nmCargo,omitempty"`
	Cbocargo  []string `xml:"CBOCargo,omitempty"`
	Nmfuncao  []string `xml:"nmFuncao,omitempty"`
	Cbofuncao []string `xml:"CBOFuncao,omitempty"`
}
