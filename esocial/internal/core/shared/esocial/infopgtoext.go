package esocial

type Infopgtoext struct {
	Paisresidext string   `xml:"paisResidExt"`
	Indnif       int64    `xml:"indNIF"`
	Nifbenef     []string `xml:"nifBenef,omitempty"`
	Frmtribut    string   `xml:"frmTribut"`
	Endext       []Endext `xml:"endExt,omitempty"`
}
