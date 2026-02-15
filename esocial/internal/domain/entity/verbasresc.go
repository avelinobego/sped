package entity

type Verbasresc struct {
	Proccs      []Proccs      `xml:"procCS,omitempty"`
	Dmdev       []Dmdev       `xml:"dmDev"`
	Procjudtrab []Procjudtrab `xml:"procJudTrab"`
	Infomv      []Infomv      `xml:"infoMV,omitempty"`
}
