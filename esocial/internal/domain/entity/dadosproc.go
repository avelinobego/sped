package entity

type Dadosproc struct {
	Indautoria   []string       `xml:"indAutoria,omitempty"`
	Indmatproc   int64          `xml:"indMatProc"`
	Observacao   []string       `xml:"observacao,omitempty"`
	Dadosprocjud []Dadosprocjud `xml:"dadosProcJud,omitempty"`
	Infosusp     []Infosusp     `xml:"infoSusp"`
}
