package esocial

type Infocategpispasep struct {
	Matricula        []string           `xml:"matricula,omitempty"`
	Codcateg         int64              `xml:"codCateg"`
	Infobasepispasep []Infobasepispasep `xml:"infoBasePisPasep"`
}
