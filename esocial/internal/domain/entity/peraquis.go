package entity

type Peraquis struct {
	Dtinicio string   `xml:"dtInicio"`
	Dtfim    []string `xml:"dtFim,omitempty"`
}
