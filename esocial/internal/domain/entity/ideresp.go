package entity

type Ideresp struct {
	Tpinsc       int64    `xml:"tpInsc"`
	Nrinsc       string   `xml:"nrInsc"`
	Dtadmrespdir []string `xml:"dtAdmRespDir,omitempty"`
	Matrespdir   []string `xml:"matRespDir,omitempty"`
}
