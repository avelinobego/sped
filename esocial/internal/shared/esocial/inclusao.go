package esocial

type Inclusao struct {
	Ideperiodo   *Ideperiodo   `xml:"idePeriodo"`
	Infocadastro *Infocadastro `xml:"infoCadastro,omitempty"`
	Ideestab     *Ideestab     `xml:"ideEstab"`
	Dadosestab   *Dadosestab   `xml:"dadosEstab"`
	Iderubrica   *Iderubrica   `xml:"ideRubrica"`
	Dadosrubrica *Dadosrubrica `xml:"dadosRubrica"`
	Idelotacao   *Idelotacao   `xml:"ideLotacao"`
	Dadoslotacao *Dadoslotacao `xml:"dadosLotacao"`
	Ideprocesso  *Ideprocesso  `xml:"ideProcesso"`
	Dadosproc    *Dadosproc    `xml:"dadosProc"`
}
