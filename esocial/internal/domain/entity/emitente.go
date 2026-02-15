package entity

type Emitente struct {
	Nmemit string   `xml:"nmEmit"`
	Ideoc  int64    `xml:"ideOC"`
	Nroc   string   `xml:"nrOC"`
	Ufoc   []string `xml:"ufOC,omitempty"`
}
