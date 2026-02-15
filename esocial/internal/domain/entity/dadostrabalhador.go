package entity

type Dadostrabalhador struct {
	Nmtrab          string            `xml:"nmTrab"`
	Sexo            string            `xml:"sexo"`
	Racacor         int64             `xml:"racaCor"`
	Estciv          []string          `xml:"estCiv,omitempty"`
	Grauinstr       string            `xml:"grauInstr"`
	Nmsoc           []string          `xml:"nmSoc,omitempty"`
	Paisnac         string            `xml:"paisNac"`
	Endereco        []string          `xml:"endereco,omitempty"`
	Trabimig        []Trabimig        `xml:"trabImig,omitempty"`
	Infodeficiencia []Infodeficiencia `xml:"infoDeficiencia,omitempty"`
	Dependente      []Dependente      `xml:"dependente"`
	Contato         []Contato         `xml:"contato,omitempty"`
}
