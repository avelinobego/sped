package esocial

type Infoemprparcial struct {
	Tpinsccontrat int64    `xml:"tpInscContrat"`
	Nrinsccontrat string   `xml:"nrInscContrat"`
	Tpinscprop    []string `xml:"tpInscProp,omitempty"`
	Nrinscprop    []string `xml:"nrInscProp,omitempty"`
}
