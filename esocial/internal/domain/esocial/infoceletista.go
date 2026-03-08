package esocial

type Infoceletista struct {
	Dtadm             string           `xml:"dtAdm"`
	Tpadmissao        int64            `xml:"tpAdmissao"`
	Indadmissao       int64            `xml:"indAdmissao"`
	Nrproctrab        []string         `xml:"nrProcTrab,omitempty"`
	Matanotjud        []string         `xml:"matAnotJud,omitempty"`
	Fgts              []Fgts           `xml:"FGTS,omitempty"`
	Tpregjor          int64            `xml:"tpRegJor"`
	Natatividade      int64            `xml:"natAtividade"`
	Dtbase            []string         `xml:"dtBase,omitempty"`
	Cnpjsindcategprof string           `xml:"cnpjSindCategProf"`
	Trabtemporario    []Trabtemporario `xml:"trabTemporario,omitempty"`
	Aprend            []Aprend         `xml:"aprend,omitempty"`
}
