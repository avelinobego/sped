package esocial

type Instpenmorte struct {
	Cpfinst      string   `xml:"cpfInst"`
	Dtinst       string   `xml:"dtInst"`
	Tpdepinst    string   `xml:"tpDepInst"`
	Descrdepinst []string `xml:"descrDepInst,omitempty"`
}
