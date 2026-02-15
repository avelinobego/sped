package entity

type Infocomplem struct {
	Nmtrab       string         `xml:"nmTrab"`
	Dtnascto     string         `xml:"dtNascto"`
	Sucessaovinc []Sucessaovinc `xml:"sucessaoVinc,omitempty"`
}
