package entity

type Exame struct {
	Dtexm         string   `xml:"dtExm"`
	Procrealizado int64    `xml:"procRealizado"`
	Obsproc       []string `xml:"obsProc,omitempty"`
	Ordexame      []string `xml:"ordExame,omitempty"`
	Indresult     []string `xml:"indResult,omitempty"`
}
