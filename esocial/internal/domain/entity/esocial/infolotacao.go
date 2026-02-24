package esocial

type Infolotacao struct {
	Inclusao  []Inclusao  `xml:"inclusao,omitempty"`
	Alteracao []Alteracao `xml:"alteracao,omitempty"`
	Exclusao  []Exclusao  `xml:"exclusao,omitempty"`
}
