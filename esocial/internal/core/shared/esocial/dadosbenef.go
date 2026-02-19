package esocial

type Dadosbenef struct {
	Nmbenefic  string       `xml:"nmBenefic"`
	Sexo       string       `xml:"sexo"`
	Racacor    int64        `xml:"racaCor"`
	Estciv     []string     `xml:"estCiv,omitempty"`
	Incfismen  string       `xml:"incFisMen"`
	Endereco   string       `xml:"endereco"`
	Dependente []Dependente `xml:"dependente"`
}
