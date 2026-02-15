package entity

type Ideproctrab struct {
	Nrproctrab  string   `xml:"nrProcTrab"`
	Cpftrab     []string `xml:"cpfTrab,omitempty"`
	Perapurpgto []string `xml:"perApurPgto,omitempty"`
	Ideseqproc  []string `xml:"ideSeqProc,omitempty"`
}
