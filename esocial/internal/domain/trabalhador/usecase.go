package trabalhador

import (
	"context"
	"fmt"
)

// Usecase defines the interface for trabalhador use cases
type Usecase interface {
	CreateTrabalhador(ctx context.Context, req *CreateTrabalhadorRequest) (*TrabalhadorResponse, error)
	GetTrabalhador(ctx context.Context, id string) (*TrabalhadorResponse, error)
	UpdateTrabalhador(ctx context.Context, id string, req *UpdateTrabalhadorRequest) (*TrabalhadorResponse, error)
	DeleteTrabalhador(ctx context.Context, id string) error
	ListTrabalhadores(ctx context.Context, req *ListRequest) (*ListResponse[TrabalhadorResponse], error)
	ValidateCPF(ctx context.Context, cpf string) error
	ValidateTrabalhadorData(ctx context.Context, req *CreateTrabalhadorRequest) error
}

// TrabalhadorUsecase implements Usecase interface
type TrabalhadorUsecase struct {
	service Service
}

// NewUsecase creates a new trabalhador usecase
func NewUsecase(service Service) Usecase {
	return &TrabalhadorUsecase{service: service}
}

// CreateTrabalhador creates a new trabalhador with validation
func (u *TrabalhadorUsecase) CreateTrabalhador(ctx context.Context, req *CreateTrabalhadorRequest) (*TrabalhadorResponse, error) {
	// Validate request data
	if err := u.ValidateTrabalhadorData(ctx, req); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	// Validate CPF
	if err := u.ValidateCPF(ctx, req.CPF); err != nil {
		return nil, fmt.Errorf("CPF validation failed: %w", err)
	}

	// Create trabalhador
	return u.service.Create(ctx, req)
}

// GetTrabalhador retrieves a trabalhador by ID
func (u *TrabalhadorUsecase) GetTrabalhador(ctx context.Context, id string) (*TrabalhadorResponse, error) {
	return u.service.GetByID(ctx, id)
}

// UpdateTrabalhador updates an existing trabalhador
func (u *TrabalhadorUsecase) UpdateTrabalhador(ctx context.Context, id string, req *UpdateTrabalhadorRequest) (*TrabalhadorResponse, error) {
	// If CPF is being updated, validate it
	if req.NIS != nil {
		if err := u.ValidateCPF(ctx, *req.NIS); err != nil {
			return nil, fmt.Errorf("CPF validation failed: %w", err)
		}
	}

	return u.service.Update(ctx, id, req)
}

// DeleteTrabalhador removes a trabalhador
func (u *TrabalhadorUsecase) DeleteTrabalhador(ctx context.Context, id string) error {
	return u.service.Delete(ctx, id)
}

// ListTrabalhadores lists trabalhadores with pagination
func (u *TrabalhadorUsecase) ListTrabalhadores(ctx context.Context, req *ListRequest) (*ListResponse[TrabalhadorResponse], error) {
	return u.service.List(ctx, req)
}

// ValidateCPF validates Brazilian CPF format and checksum
func (u *TrabalhadorUsecase) ValidateCPF(ctx context.Context, cpf string) error {
	if len(cpf) != 11 {
		return fmt.Errorf("CPF must have 11 digits")
	}

	// Check if all digits are the same (invalid CPF)
	allSame := true
	for i := 1; i < len(cpf); i++ {
		if cpf[i] != cpf[0] {
			allSame = false
			break
		}
	}
	if allSame {
		return fmt.Errorf("CPF is invalid")
	}

	// Calculate first verification digit
	sum := 0
	for i := 0; i < 9; i++ {
		digit := int(cpf[i] - '0')
		sum += digit * (10 - i)
	}
	remainder := sum % 11
	firstDigit := 0
	if remainder >= 2 {
		firstDigit = 11 - remainder
	}

	// Check first verification digit
	if firstDigit != int(cpf[9]-'0') {
		return fmt.Errorf("CPF is invalid")
	}

	// Calculate second verification digit
	sum = 0
	for i := 0; i < 10; i++ {
		digit := int(cpf[i] - '0')
		sum += digit * (11 - i)
	}
	remainder = sum % 11
	secondDigit := 0
	if remainder >= 2 {
		secondDigit = 11 - remainder
	}

	// Check second verification digit
	if secondDigit != int(cpf[10]-'0') {
		return fmt.Errorf("CPF is invalid")
	}

	return nil
}

// ValidateTrabalhadorData validates trabalhador data according to eSocial rules
func (u *TrabalhadorUsecase) ValidateTrabalhadorData(ctx context.Context, req *CreateTrabalhadorRequest) error {
	// Validate required fields
	if req.EmpresaID == "" {
		return fmt.Errorf("empresa_id is required")
	}

	if req.CPF == "" {
		return fmt.Errorf("CPF is required")
	}

	if req.Nome == "" {
		return fmt.Errorf("nome is required")
	}

	if req.PaisNascimento == "" {
		return fmt.Errorf("pais_nascimento is required")
	}

	if req.PaisNacionalidade == "" {
		return fmt.Errorf("pais_nacionalidade is required")
	}

	// Validate sex
	if req.Sexo != "M" && req.Sexo != "F" {
		return fmt.Errorf("sexo must be M or F")
	}

	// Validate race/color if provided
	if req.RacaCor != nil {
		validRacaCor := []string{"1", "2", "3", "4", "5", "6"}
		valid := false
		for _, v := range validRacaCor {
			if *req.RacaCor == v {
				valid = true
				break
			}
		}
		if !valid {
			return fmt.Errorf("raca_cor must be one of: 1, 2, 3, 4, 5, 6")
		}
	}

	// Validate marital status if provided
	if req.EstadoCivil != nil {
		validEstadoCivil := []string{"1", "2", "3", "4", "5"}
		valid := false
		for _, v := range validEstadoCivil {
			if *req.EstadoCivil == v {
				valid = true
				break
			}
		}
		if !valid {
			return fmt.Errorf("estado_civil must be one of: 1, 2, 3, 4, 5")
		}
	}

	// Validate education level if provided
	if req.GrauInstrucao != nil {
		validGrauInstrucao := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12"}
		valid := false
		for _, v := range validGrauInstrucao {
			if *req.GrauInstrucao == v {
				valid = true
				break
			}
		}
		if !valid {
			return fmt.Errorf("grau_instrucao must be one of: 01, 02, 03, 04, 05, 06, 07, 08, 09, 10, 11, 12")
		}
	}

	// Validate disability type if disability is indicated
	if req.PossuiDeficiencia && req.TipoDeficiencia == nil {
		return fmt.Errorf("tipo_deficiencia is required when possui_deficiencia is true")
	}

	if req.TipoDeficiencia != nil {
		validTipoDeficiencia := []string{"1", "2", "3", "4", "5", "6"}
		valid := false
		for _, v := range validTipoDeficiencia {
			if *req.TipoDeficiencia == v {
				valid = true
				break
			}
		}
		if !valid {
			return fmt.Errorf("tipo_deficiencia must be one of: 1, 2, 3, 4, 5, 6")
		}
	}

	// Validate driver's license category if provided
	if req.CategoriaCNH != nil {
		validCategorias := []string{"A", "B", "C", "D", "E", "AB", "AC", "AD", "AE"}
		valid := false
		for _, v := range validCategorias {
			if *req.CategoriaCNH == v {
				valid = true
				break
			}
		}
		if !valid {
			return fmt.Errorf("categoria_cnh must be one of: A, B, C, D, E, AB, AC, AD, AE")
		}
	}

	// Validate account type if provided
	if req.TipoConta != nil {
		if *req.TipoConta != "1" && *req.TipoConta != "2" {
			return fmt.Errorf("tipo_conta must be 1 or 2")
		}
	}

	return nil
}