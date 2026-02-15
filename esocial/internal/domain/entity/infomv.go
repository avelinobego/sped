package entity

type Infomv struct {
	Indmv         int64           `xml:"indMV"`
	Remunoutrempr []Remunoutrempr `xml:"remunOutrEmpr"`
}
