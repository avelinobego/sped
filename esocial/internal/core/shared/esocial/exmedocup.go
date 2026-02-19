package esocial

type Exmedocup struct {
	Tpexameocup int64       `xml:"tpExameOcup"`
	Aso         Aso         `xml:"aso"`
	Respmonit   []Respmonit `xml:"respMonit,omitempty"`
}
