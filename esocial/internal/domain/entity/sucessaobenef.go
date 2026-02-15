package entity

type Sucessaobenef struct {
	Cnpjorgaoant   string   `xml:"cnpjOrgaoAnt"`
	Nrbeneficioant string   `xml:"nrBeneficioAnt"`
	Dttransf       string   `xml:"dtTransf"`
	Observacao     []string `xml:"observacao,omitempty"`
}
