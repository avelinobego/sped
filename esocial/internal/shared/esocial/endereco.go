package esocial

type Endereco struct {
	Brasil   []Brasil   `xml:"brasil,omitempty"`
	Exterior []Exterior `xml:"exterior,omitempty"`
}
