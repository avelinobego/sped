package entity

type Dadoslotacao struct {
	Tplotacao       string            `xml:"tpLotacao"`
	Tpinsc          []string          `xml:"tpInsc,omitempty"`
	Nrinsc          []string          `xml:"nrInsc,omitempty"`
	Fpaslotacao     Fpaslotacao       `xml:"fpasLotacao"`
	Infoemprparcial []Infoemprparcial `xml:"infoEmprParcial,omitempty"`
	Dadosopport     []Dadosopport     `xml:"dadosOpPort,omitempty"`
}
