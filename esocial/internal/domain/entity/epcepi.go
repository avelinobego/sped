package entity

type Epcepi struct {
	Utilizepc int64      `xml:"utilizEPC"`
	Eficepc   []string   `xml:"eficEpc,omitempty"`
	Utilizepi int64      `xml:"utilizEPI"`
	Eficepi   []string   `xml:"eficEpi,omitempty"`
	Epi       []Epi      `xml:"epi"`
	Epicompl  []Epicompl `xml:"epiCompl,omitempty"`
}
