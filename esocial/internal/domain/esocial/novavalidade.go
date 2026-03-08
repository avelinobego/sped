package esocial

type Novavalidade struct {
	Inivalid string   `xml:"iniValid"`
	Fimvalid []string `xml:"fimValid,omitempty"`
}
