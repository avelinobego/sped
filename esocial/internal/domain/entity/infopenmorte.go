package entity

type Infopenmorte struct {
	Tppenmorte   int64          `xml:"tpPenMorte"`
	Instpenmorte []Instpenmorte `xml:"instPenMorte,omitempty"`
}
