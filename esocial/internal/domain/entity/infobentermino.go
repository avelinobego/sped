package entity

type Infobentermino struct {
	Dttermbeneficio string   `xml:"dtTermBeneficio"`
	Mtvtermino      string   `xml:"mtvTermino"`
	Cnpjorgaosuc    []string `xml:"cnpjOrgaoSuc,omitempty"`
	Novocpf         []string `xml:"novoCPF,omitempty"`
}
