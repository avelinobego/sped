package entity

type Desligamento struct {
	Mtvdeslig    string   `xml:"mtvDeslig"`
	Dtdeslig     string   `xml:"dtDeslig"`
	Dtprojfimapi []string `xml:"dtProjFimAPI,omitempty"`
}
