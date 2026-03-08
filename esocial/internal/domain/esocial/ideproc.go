package esocial

type Ideproc struct {
	Ideseqproc   []string       `xml:"ideSeqProc,omitempty"`
	Obs          []string       `xml:"obs,omitempty"`
	Perapurpgto  string         `xml:"perApurPgto"`
	Perapur      string         `xml:"perApur"`
	Infotributos []Infotributos `xml:"infoTributos"`
	Infocrirrf   []Infocrirrf   `xml:"infoCRIRRF"`
	Origem       int64          `xml:"origem"`
	Nrproctrab   string         `xml:"nrProcTrab"`
}
