package esocial

type Infodirigentesindical struct {
	Categorig  int64    `xml:"categOrig"`
	Tpinsc     []string `xml:"tpInsc,omitempty"`
	Nrinsc     []string `xml:"nrInsc,omitempty"`
	Dtadmorig  []string `xml:"dtAdmOrig,omitempty"`
	Matricorig []string `xml:"matricOrig,omitempty"`
	Tpregtrab  []string `xml:"tpRegTrab,omitempty"`
	Tpregprev  int64    `xml:"tpRegPrev"`
}
