package entity

type Infoestab struct {
	Inclusao      []Inclusao      `xml:"inclusao,omitempty"`
	Alteracao     []Alteracao     `xml:"alteracao,omitempty"`
	Exclusao      []Exclusao      `xml:"exclusao,omitempty"`
	Cnaeprep      int64           `xml:"cnaePrep"`
	Cnpjresp      []string        `xml:"cnpjResp,omitempty"`
	Aliqrat       int64           `xml:"aliqRat"`
	Fap           []string        `xml:"fap,omitempty"`
	Aliqratajust  []string        `xml:"aliqRatAjust,omitempty"`
	Infoestabref  []Infoestabref  `xml:"infoEstabRef,omitempty"`
	Infocomplobra []Infocomplobra `xml:"infoComplObra,omitempty"`
}
