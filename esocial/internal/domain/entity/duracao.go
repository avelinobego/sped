package entity

type Duracao struct {
	Tpcontr   int64    `xml:"tpContr"`
	Dtterm    []string `xml:"dtTerm,omitempty"`
	Clauassec []string `xml:"clauAssec,omitempty"`
	Objdet    []string `xml:"objDet,omitempty"`
}
