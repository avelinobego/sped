package entity

type Infocontrib struct {
	Classtrib string   `xml:"classTrib"`
	Infopj    []Infopj `xml:"infoPJ,omitempty"`
}
