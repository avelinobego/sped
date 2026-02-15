package entity

type Endext struct {
	Enddsclograd []string `xml:"endDscLograd,omitempty"`
	Endnrlograd  []string `xml:"endNrLograd,omitempty"`
	Endcomplem   []string `xml:"endComplem,omitempty"`
	Endbairro    []string `xml:"endBairro,omitempty"`
	Endcidade    []string `xml:"endCidade,omitempty"`
	Endestado    []string `xml:"endEstado,omitempty"`
	Endcodpostal []string `xml:"endCodPostal,omitempty"`
	Telef        []string `xml:"telef,omitempty"`
}
