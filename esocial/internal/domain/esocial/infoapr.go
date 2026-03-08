package esocial

type Infoapr struct {
	Nrprocjud   []string      `xml:"nrProcJud,omitempty"`
	Infoenteduc []Infoenteduc `xml:"infoEntEduc"`
}
