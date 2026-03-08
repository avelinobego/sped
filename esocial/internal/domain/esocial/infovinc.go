package esocial

type Infovinc struct {
	Tpregtrab    int64          `xml:"tpRegTrab"`
	Tpregprev    int64          `xml:"tpRegPrev"`
	Dtadm        string         `xml:"dtAdm"`
	Tmpparc      []string       `xml:"tmpParc,omitempty"`
	Duracao      []Duracao      `xml:"duracao,omitempty"`
	Observacoes  []Observacoes  `xml:"observacoes"`
	Sucessaovinc []Sucessaovinc `xml:"sucessaoVinc,omitempty"`
	Infodeslig   []Infodeslig   `xml:"infoDeslig,omitempty"`
}
