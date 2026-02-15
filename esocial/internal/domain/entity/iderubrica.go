package entity

type Iderubrica struct {
	Codrubr    string   `xml:"codRubr"`
	Idetabrubr string   `xml:"ideTabRubr"`
	Inivalid   string   `xml:"iniValid"`
	Fimvalid   []string `xml:"fimValid,omitempty"`
}
