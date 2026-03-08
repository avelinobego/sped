package esocial

type Sucessaovinc struct {
	Cnpjorgaoant string   `xml:"cnpjOrgaoAnt"`
	Dtexercicio  string   `xml:"dtExercicio"`
	Observacao   []string `xml:"observacao,omitempty"`
	Dttransf     string   `xml:"dtTransf"`
	Tpinsc       int64    `xml:"tpInsc"`
	Nrinsc       string   `xml:"nrInsc"`
	Matricant    []string `xml:"matricAnt,omitempty"`
	Dtadm        string   `xml:"dtAdm"`
}
