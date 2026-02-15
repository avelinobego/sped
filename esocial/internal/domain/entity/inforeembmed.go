package entity

type Inforeembmed struct {
	Indorgreemb  int64          `xml:"indOrgReemb"`
	Cnpjoper     []string       `xml:"cnpjOper,omitempty"`
	Regans       []string       `xml:"regANS,omitempty"`
	Detreembtit  []Detreembtit  `xml:"detReembTit"`
	Inforeembdep []Inforeembdep `xml:"infoReembDep"`
}
