package entity

type Infoprocesso struct {
	Inclusao    []Inclusao  `xml:"inclusao,omitempty"`
	Alteracao   []Alteracao `xml:"alteracao,omitempty"`
	Exclusao    []Exclusao  `xml:"exclusao,omitempty"`
	Origem      int64       `xml:"origem"`
	Obsproctrab []string    `xml:"obsProcTrab,omitempty"`
	Dadoscompl  string      `xml:"dadosCompl"`
	Nrproctrab  string      `xml:"nrProcTrab"`
	Dtsent      string      `xml:"dtSent"`
	Ufvara      string      `xml:"ufVara"`
	Codmunic    int64       `xml:"codMunic"`
	Idvara      int64       `xml:"idVara"`
}
