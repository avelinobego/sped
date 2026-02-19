package esocial

type Mudancacpf struct {
	Cpfant         string   `xml:"cpfAnt"`
	Nrbeneficioant string   `xml:"nrBeneficioAnt"`
	Dtaltcpf       string   `xml:"dtAltCPF"`
	Observacao     []string `xml:"observacao,omitempty"`
}
