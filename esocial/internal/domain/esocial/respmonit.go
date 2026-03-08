package esocial

type Respmonit struct {
	Cpfresp []string `xml:"cpfResp,omitempty"`
	Nmresp  string   `xml:"nmResp"`
	Nrcrm   string   `xml:"nrCRM"`
	Ufcrm   string   `xml:"ufCRM"`
}
