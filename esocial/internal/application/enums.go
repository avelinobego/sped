// Package application provides high-level services for the eSocial system
package application

// RacaCor represents the worker's race/color codes according to eSocial standards
type RacaCor int64

const (
	RacaCorBranca       RacaCor = 1 // Branca (White)
	RacaCorPreta        RacaCor = 2 // Preta (Black)
	RacaCorParda        RacaCor = 3 // Parda (Brown)
	RacaCorAmarela      RacaCor = 4 // Amarela (Yellow)
	RacaCorIndígena     RacaCor = 5 // Indígena (Indigenous)
	RacaCorNãoInformado RacaCor = 6 // Não informado (Not informed)
)

// String returns the string representation of RacaCor
func (r RacaCor) String() string {
	switch r {
	case RacaCorBranca:
		return "Branca"
	case RacaCorPreta:
		return "Preta"
	case RacaCorParda:
		return "Parda"
	case RacaCorAmarela:
		return "Amarela"
	case RacaCorIndígena:
		return "Indígena"
	case RacaCorNãoInformado:
		return "Não informado"
	default:
		return "Não informado"
	}
}

// ParseRacaCor converts a string code to RacaCor enumeration
func ParseRacaCor(code string) RacaCor {
	switch code {
	case "1":
		return RacaCorBranca
	case "2":
		return RacaCorPreta
	case "3":
		return RacaCorParda
	case "4":
		return RacaCorAmarela
	case "5":
		return RacaCorIndígena
	case "6":
		return RacaCorNãoInformado
	default:
		return RacaCorNãoInformado
	}
}

// ToInt64 converts RacaCor to int64 for XML serialization
func (r RacaCor) ToInt64() int64 {
	return int64(r)
}
