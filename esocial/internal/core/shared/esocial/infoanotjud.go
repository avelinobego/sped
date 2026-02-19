package esocial

type Infoanotjud struct {
	Cpftrab       string         `xml:"cpfTrab"`
	Nmtrab        string         `xml:"nmTrab"`
	Dtnascto      string         `xml:"dtNascto"`
	Dtadm         string         `xml:"dtAdm"`
	Matricula     string         `xml:"matricula"`
	Codcateg      int64          `xml:"codCateg"`
	Natatividade  int64          `xml:"natAtividade"`
	Tpcontr       int64          `xml:"tpContr"`
	Dtterm        []string       `xml:"dtTerm,omitempty"`
	Tpinsctrab    []string       `xml:"tpInscTrab,omitempty"`
	Localtrabalho []string       `xml:"localTrabalho,omitempty"`
	Tpregtrab     int64          `xml:"tpRegTrab"`
	Tpregprev     int64          `xml:"tpRegPrev"`
	Cargo         []Cargo        `xml:"cargo"`
	Remuneracao   []Remuneracao  `xml:"remuneracao"`
	Incorporacao  []Incorporacao `xml:"incorporacao"`
	Afastamento   []Afastamento  `xml:"afastamento,omitempty"`
	Desligamento  []Desligamento `xml:"desligamento,omitempty"`
}
