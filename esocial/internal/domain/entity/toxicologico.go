package entity

type Toxicologico struct {
	Dtexame     string   `xml:"dtExame"`
	Cnpjlab     string   `xml:"cnpjLab"`
	Codseqexame string   `xml:"codSeqExame"`
	Nmmed       string   `xml:"nmMed"`
	Nrcrm       []string `xml:"nrCRM,omitempty"`
	Ufcrm       []string `xml:"ufCRM,omitempty"`
}
