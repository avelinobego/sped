package entity

type Beneficiario struct {
	Nmbenefic   string       `xml:"nmBenefic"`
	Dtnascto    string       `xml:"dtNascto"`
	Dtinicio    string       `xml:"dtInicio"`
	Sexo        []string     `xml:"sexo,omitempty"`
	Racacor     int64        `xml:"racaCor"`
	Estciv      []string     `xml:"estCiv,omitempty"`
	Incfismen   string       `xml:"incFisMen"`
	Dtincfismen []string     `xml:"dtIncFisMen,omitempty"`
	Endereco    string       `xml:"endereco"`
	Dependente  []Dependente `xml:"dependente"`
	Cpfbenef    string       `xml:"cpfBenef"`
	Matricula   []string     `xml:"matricula,omitempty"`
	Cnpjorigem  []string     `xml:"cnpjOrigem,omitempty"`
}
