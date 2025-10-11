package tipos

type IdeTrabalhador struct {
	Cpf     CPF    `xml:"cpfTrab"`
	NisTrab uint64 `xml:"nisTrab"`
	Nome    string `xml:"nmTrab"`
}
