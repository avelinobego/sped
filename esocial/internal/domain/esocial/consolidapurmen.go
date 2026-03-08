package esocial

type Consolidapurmen struct {
	Crmen              string   `xml:"CRMen"`
	Vlrrendtrib        float64  `xml:"vlrRendTrib"`
	Vlrrendtrib13      float64  `xml:"vlrRendTrib13"`
	Vlrprevoficial     float64  `xml:"vlrPrevOficial"`
	Vlrprevoficial13   float64  `xml:"vlrPrevOficial13"`
	Vlrcrmen           float64  `xml:"vlrCRMen"`
	Vlrcr13men         float64  `xml:"vlrCR13Men"`
	Vlrparcisenta65    float64  `xml:"vlrParcIsenta65"`
	Vlrparcisenta65dec float64  `xml:"vlrParcIsenta65Dec"`
	Vlrdiarias         float64  `xml:"vlrDiarias"`
	Vlrajudacusto      float64  `xml:"vlrAjudaCusto"`
	Vlrindrescontrato  float64  `xml:"vlrIndResContrato"`
	Vlrabonopec        float64  `xml:"vlrAbonoPec"`
	Vlrrendmolegrave   float64  `xml:"vlrRendMoleGrave"`
	Vlrrendmolegrave13 float64  `xml:"vlrRendMoleGrave13"`
	Vlrauxmoradia      float64  `xml:"vlrAuxMoradia"`
	Vlrbolsamedico     float64  `xml:"vlrBolsaMedico"`
	Vlrbolsamedico13   float64  `xml:"vlrBolsaMedico13"`
	Vlrjurosmora       float64  `xml:"vlrJurosMora"`
	Vlrisenoutros      float64  `xml:"vlrIsenOutros"`
	Descrendimento     []string `xml:"descRendimento,omitempty"`
}
