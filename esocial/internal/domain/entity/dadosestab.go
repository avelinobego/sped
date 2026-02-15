package entity

type Dadosestab struct {
	Cnaeprep   int64        `xml:"cnaePrep"`
	Cnpjresp   []string     `xml:"cnpjResp,omitempty"`
	Aliqgilrat []Aliqgilrat `xml:"aliqGilrat,omitempty"`
	Infocaepf  []Infocaepf  `xml:"infoCaepf,omitempty"`
	Infoobra   []Infoobra   `xml:"infoObra,omitempty"`
	Infotrab   []Infotrab   `xml:"infoTrab,omitempty"`
}
