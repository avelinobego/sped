package esocial

type Infomandelet struct {
	Cnpjmandelet  string   `xml:"cnpjMandElet"`
	Categorig     int64    `xml:"categOrig"`
	Cnpjorig      string   `xml:"cnpjOrig"`
	Matricorig    string   `xml:"matricOrig"`
	Dtexercorig   string   `xml:"dtExercOrig"`
	Tpregtrab     int64    `xml:"tpRegTrab"`
	Indremuncargo []string `xml:"indRemunCargo,omitempty"`
	Tpregprev     int64    `xml:"tpRegPrev"`
}
