package esocial

type Infocategincid struct {
	Matricula  []string     `xml:"matricula,omitempty"`
	Codcateg   int64        `xml:"codCateg"`
	Indsimples []string     `xml:"indSimples,omitempty"`
	Infobasecs []Infobasecs `xml:"infoBaseCS"`
	Calcterc   []Calcterc   `xml:"calcTerc"`
	Infoperref []Infoperref `xml:"infoPerRef"`
}
