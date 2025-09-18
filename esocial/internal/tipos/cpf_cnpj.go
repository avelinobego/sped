package tipos

type CpfCnpjConstraint interface {
	CPF | CNPJ
	Validar() error
	Tipo() int
}
