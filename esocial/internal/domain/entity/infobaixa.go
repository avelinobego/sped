package entity

type Infobaixa struct {
	Mtvdeslig    string   `xml:"mtvDeslig"`
	Dtdeslig     string   `xml:"dtDeslig"`
	Dtprojfimapi []string `xml:"dtProjFimAPI,omitempty"`
	Nrproctrab   string   `xml:"nrProcTrab"`
	Observacao   []string `xml:"observacao,omitempty"`
}
