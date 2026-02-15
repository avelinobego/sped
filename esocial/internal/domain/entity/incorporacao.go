package entity

type Incorporacao struct {
	Tpinsc    []string `xml:"tpInsc,omitempty"`
	Nrinsc    []string `xml:"nrInsc,omitempty"`
	Matincorp string   `xml:"matIncorp"`
}
