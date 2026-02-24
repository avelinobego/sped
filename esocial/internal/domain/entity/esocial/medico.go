package esocial

type Medico struct {
	Nmmed string   `xml:"nmMed"`
	Nrcrm []string `xml:"nrCRM,omitempty"`
	Ufcrm []string `xml:"ufCRM,omitempty"`
}
