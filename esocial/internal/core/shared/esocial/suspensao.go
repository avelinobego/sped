package esocial

type Suspensao struct {
	Mtvsuspensao string   `xml:"mtvSuspensao"`
	Dscsuspensao []string `xml:"dscSuspensao,omitempty"`
}
