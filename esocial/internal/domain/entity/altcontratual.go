package entity

type Altcontratual struct {
	Dtalteracao string   `xml:"dtAlteracao"`
	Dtef        []string `xml:"dtEf,omitempty"`
	Dscalt      []string `xml:"dscAlt,omitempty"`
	Vinculo     Vinculo  `xml:"vinculo"`
}
