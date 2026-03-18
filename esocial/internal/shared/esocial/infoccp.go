package esocial

type Infoccp struct {
	Dtccp   string   `xml:"dtCCP"`
	Tpccp   int64    `xml:"tpCCP"`
	Cnpjccp []string `xml:"cnpjCCP,omitempty"`
}
