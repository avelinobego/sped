package entity

type Idelotacao struct {
	Inivalid string `xml:"iniValid"`
	Fimvalid string `xml:"fimValid,omitempty"`
}
