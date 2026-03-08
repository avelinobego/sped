package esocial

type Infoafastamento struct {
	Iniafastamento []Iniafastamento `xml:"iniAfastamento,omitempty"`
	Inforetif      []Inforetif      `xml:"infoRetif,omitempty"`
	Fimafastamento []Fimafastamento `xml:"fimAfastamento,omitempty"`
}
