package esocial

type Infodeslig struct {
	Dtavprv         []string          `xml:"dtAvPrv,omitempty"`
	Indpagtoapi     string            `xml:"indPagtoAPI"`
	Nrproctrab      []string          `xml:"nrProcTrab,omitempty"`
	Indpdv          []string          `xml:"indPDV,omitempty"`
	Infointerm      []Infointerm      `xml:"infoInterm"`
	Observacoes     []Observacoes     `xml:"observacoes"`
	Sucessaovinc    []Sucessaovinc    `xml:"sucessaoVinc,omitempty"`
	Transftit       []Transftit       `xml:"transfTit,omitempty"`
	Mudancacpf      []Mudancacpf      `xml:"mudancaCPF,omitempty"`
	Verbasresc      []Verbasresc      `xml:"verbasResc,omitempty"`
	Remunaposdeslig []Remunaposdeslig `xml:"remunAposDeslig,omitempty"`
	Consigfgts      []Consigfgts      `xml:"consigFGTS"`
	Dtdeslig        string            `xml:"dtDeslig"`
	Mtvdeslig       string            `xml:"mtvDeslig"`
	Dtprojfimapi    []string          `xml:"dtProjFimAPI,omitempty"`
	Pensalim        []string          `xml:"pensAlim,omitempty"`
	Percaliment     []string          `xml:"percAliment,omitempty"`
	Vralim          []string          `xml:"vrAlim,omitempty"`
}
