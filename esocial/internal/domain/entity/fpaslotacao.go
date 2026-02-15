package entity

type Fpaslotacao struct {
	Fpas                 int64                  `xml:"fpas"`
	Codtercs             string                 `xml:"codTercs"`
	Codtercssusp         []string               `xml:"codTercsSusp,omitempty"`
	Infoprocjudterceiros []Infoprocjudterceiros `xml:"infoProcJudTerceiros,omitempty"`
}
