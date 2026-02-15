package entity

type Basecalculo struct {
	Vrbccpmensal float64        `xml:"vrBcCpMensal"`
	Vrbccp13     []string       `xml:"vrBcCp13,omitempty"`
	Infoagnocivo []Infoagnocivo `xml:"infoAgNocivo,omitempty"`
}
