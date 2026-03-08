package esocial

type Infocomplementares struct {
	Fgts                  []Fgts                  `xml:"FGTS,omitempty"`
	Cargofuncao           []Cargofuncao           `xml:"cargoFuncao,omitempty"`
	Remuneracao           []Remuneracao           `xml:"remuneracao,omitempty"`
	Infodirigentesindical []Infodirigentesindical `xml:"infoDirigenteSindical,omitempty"`
	Infotrabcedido        []Infotrabcedido        `xml:"infoTrabCedido,omitempty"`
	Infomandelet          []Infomandelet          `xml:"infoMandElet,omitempty"`
	Infoestagiario        []Infoestagiario        `xml:"infoEstagiario,omitempty"`
	Localtrabgeral        []Localtrabgeral        `xml:"localTrabGeral,omitempty"`
}
