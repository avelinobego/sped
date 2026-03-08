package esocial

type Infoestatutario struct {
	Tpprov       int64    `xml:"tpProv"`
	Dtexercicio  string   `xml:"dtExercicio"`
	Dtiniabono   []string `xml:"dtIniAbono,omitempty"`
	Tpplanrp     int64    `xml:"tpPlanRP"`
	Indtetorgps  string   `xml:"indTetoRGPS"`
	Indabonoperm string   `xml:"indAbonoPerm"`
}
