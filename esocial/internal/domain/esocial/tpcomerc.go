package esocial

type Tpcomerc struct {
	Indcomerc   int64         `xml:"indComerc"`
	Vrtotcom    float64       `xml:"vrTotCom"`
	Ideadquir   []Ideadquir   `xml:"ideAdquir"`
	Infoprocjud []Infoprocjud `xml:"infoProcJud"`
}
