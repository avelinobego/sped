package esocial

type Basesaquis struct {
	Indaquis    int64   `xml:"indAquis"`
	Vlraquis    float64 `xml:"vlrAquis"`
	Vrcpdescpr  float64 `xml:"vrCPDescPR"`
	Vrcpnret    float64 `xml:"vrCPNRet"`
	Vrratnret   float64 `xml:"vrRatNRet"`
	Vrsenarnret float64 `xml:"vrSenarNRet"`
	Vrcpcalcpr  float64 `xml:"vrCPCalcPR"`
	Vrratdescpr float64 `xml:"vrRatDescPR"`
	Vrratcalcpr float64 `xml:"vrRatCalcPR"`
	Vrsenardesc float64 `xml:"vrSenarDesc"`
	Vrsenarcalc float64 `xml:"vrSenarCalc"`
}
