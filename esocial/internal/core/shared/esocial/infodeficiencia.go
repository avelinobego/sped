package esocial

type Infodeficiencia struct {
	Deffisica      string   `xml:"defFisica"`
	Defvisual      string   `xml:"defVisual"`
	Defauditiva    string   `xml:"defAuditiva"`
	Defmental      string   `xml:"defMental"`
	Defintelectual string   `xml:"defIntelectual"`
	Reabreadap     string   `xml:"reabReadap"`
	Observacao     []string `xml:"observacao,omitempty"`
}
