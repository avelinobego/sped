package trabalhador

import (
	"context"
	"errors"
	"testing"
	"time"
)

type mockService struct {
	createErr error
	getErr    error
	getCPFErr error
	updateErr error
	deleteErr error
	listErr   error
	responses map[string]*TrabalhadorResponse
	listings  []*TrabalhadorResponse
}

func (m *mockService) Create(ctx context.Context, req *CreateTrabalhadorRequest) (*TrabalhadorResponse, error) {
	if m.createErr != nil {
		return nil, m.createErr
	}
	return &TrabalhadorResponse{
		ID:              "t1",
		EmpresaID:       req.EmpresaID,
		CPF:             req.CPF,
		Nome:            req.Nome,
		Sexo:            req.Sexo,
		DataNascimento:  req.DataNascimento,
		DataCadastro:    time.Now(),
		DataAtualizacao: time.Now(),
		Situacao:        "ATIVO",
	}, nil
}

func (m *mockService) GetByID(ctx context.Context, id string) (*TrabalhadorResponse, error) {
	if m.getErr != nil {
		return nil, m.getErr
	}
	if m.responses != nil {
		if resp, ok := m.responses[id]; ok {
			return resp, nil
		}
	}
	return nil, nil
}

func (m *mockService) GetByCPF(ctx context.Context, cpf string) (*TrabalhadorResponse, error) {
	if m.getCPFErr != nil {
		return nil, m.getCPFErr
	}
	return nil, nil
}

func (m *mockService) Update(ctx context.Context, id string, req *UpdateTrabalhadorRequest) (*TrabalhadorResponse, error) {
	if m.updateErr != nil {
		return nil, m.updateErr
	}
	return &TrabalhadorResponse{
		ID:              "t1",
		DataAtualizacao: time.Now(),
	}, nil
}

func (m *mockService) Delete(ctx context.Context, id string) error {
	return m.deleteErr
}

func (m *mockService) List(ctx context.Context, req *ListRequest) (*ListResponse[TrabalhadorResponse], error) {
	if m.listErr != nil {
		return nil, m.listErr
	}
	data := make([]TrabalhadorResponse, len(m.listings))
	for i, v := range m.listings {
		data[i] = *v
	}
	return &ListResponse[TrabalhadorResponse]{
		Total:      int64(len(m.listings)),
		Page:       1,
		Limit:      10,
		TotalPages: 1,
		Data:       data,
	}, nil
}

func setupUsecase() (*TrabalhadorUsecase, *mockService) {
	mockSvc := &mockService{
		responses: map[string]*TrabalhadorResponse{
			"t1": {
				ID:              "t1",
				EmpresaID:       "e1",
				CPF:             "11144477735",
				Nome:            "João Silva",
				Sexo:            "M",
				DataNascimento:  time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
				DataCadastro:    time.Now(),
				DataAtualizacao: time.Now(),
				Situacao:        "ATIVO",
			},
		},
		listings: make([]*TrabalhadorResponse, 0),
	}
	return &TrabalhadorUsecase{service: mockSvc}, mockSvc
}

// TestUsecaseCreateTrabalhador tests create operations
func TestUsecaseCreateTrabalhador(t *testing.T) {
	ctx := context.Background()
	u, mockSvc := setupUsecase()

	t.Run("valid create request", func(t *testing.T) {
		req := &CreateTrabalhadorRequest{
			EmpresaID:         "e1",
			CPF:               "11144477735",
			Nome:              "João Silva",
			DataNascimento:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Sexo:              "M",
			PaisNascimento:    "BRA",
			PaisNacionalidade: "BRA",
		}
		resp, err := u.CreateTrabalhador(ctx, req)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if resp == nil {
			t.Fatal("expected response, got nil")
		}
		if resp.CPF != "11144477735" {
			t.Fatalf("expected CPF 11144477735, got %s", resp.CPF)
		}
	})

	t.Run("invalid cpf format", func(t *testing.T) {
		req := &CreateTrabalhadorRequest{
			EmpresaID:         "e1",
			CPF:               "123",
			Nome:              "João Silva",
			DataNascimento:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Sexo:              "M",
			PaisNascimento:    "BRA",
			PaisNacionalidade: "BRA",
		}
		_, err := u.CreateTrabalhador(ctx, req)
		if err == nil {
			t.Fatal("expected error for invalid CPF format")
		}
	})

	t.Run("missing empresa_id", func(t *testing.T) {
		req := &CreateTrabalhadorRequest{
			CPF:               "12345678901",
			Nome:              "João Silva",
			DataNascimento:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Sexo:              "M",
			PaisNascimento:    "BRA",
			PaisNacionalidade: "BRA",
		}
		_, err := u.CreateTrabalhador(ctx, req)
		if err == nil {
			t.Fatal("expected error for missing empresa_id")
		}
	})

	t.Run("missing nome", func(t *testing.T) {
		req := &CreateTrabalhadorRequest{
			EmpresaID:         "e1",
			CPF:               "12345678901",
			DataNascimento:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Sexo:              "M",
			PaisNascimento:    "BRA",
			PaisNacionalidade: "BRA",
		}
		_, err := u.CreateTrabalhador(ctx, req)
		if err == nil {
			t.Fatal("expected error for missing nome")
		}
	})

	t.Run("invalid sexo", func(t *testing.T) {
		req := &CreateTrabalhadorRequest{
			EmpresaID:         "e1",
			CPF:               "12345678901",
			Nome:              "João Silva",
			DataNascimento:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Sexo:              "X",
			PaisNascimento:    "BRA",
			PaisNacionalidade: "BRA",
		}
		_, err := u.CreateTrabalhador(ctx, req)
		if err == nil {
			t.Fatal("expected error for invalid sexo")
		}
	})

	t.Run("raca_cor validation", func(t *testing.T) {
		racaCor := "9"
		req := &CreateTrabalhadorRequest{
			EmpresaID:         "e1",
			CPF:               "12345678901",
			Nome:              "João Silva",
			DataNascimento:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Sexo:              "M",
			RacaCor:           &racaCor,
			PaisNascimento:    "BRA",
			PaisNacionalidade: "BRA",
		}
		_, err := u.CreateTrabalhador(ctx, req)
		if err == nil {
			t.Fatal("expected error for invalid raca_cor")
		}
	})

	t.Run("valid raca_cor", func(t *testing.T) {
		racaCor := "1"
		req := &CreateTrabalhadorRequest{
			EmpresaID:         "e1",
			CPF:               "11144477735",
			Nome:              "João Silva",
			DataNascimento:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Sexo:              "M",
			RacaCor:           &racaCor,
			PaisNascimento:    "BRA",
			PaisNacionalidade: "BRA",
		}
		resp, err := u.CreateTrabalhador(ctx, req)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if resp == nil {
			t.Fatal("expected response, got nil")
		}
	})

	t.Run("tipo_deficiencia required when possui_deficiencia true", func(t *testing.T) {
		req := &CreateTrabalhadorRequest{
			EmpresaID:         "e1",
			CPF:               "12345678901",
			Nome:              "João Silva",
			DataNascimento:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Sexo:              "M",
			PaisNascimento:    "BRA",
			PaisNacionalidade: "BRA",
			PossuiDeficiencia: true,
		}
		_, err := u.CreateTrabalhador(ctx, req)
		if err == nil {
			t.Fatal("expected error for missing tipo_deficiencia")
		}
	})

	t.Run("invalid categoria_cnh", func(t *testing.T) {
		categoriaCNH := "Z"
		req := &CreateTrabalhadorRequest{
			EmpresaID:         "e1",
			CPF:               "12345678901",
			Nome:              "João Silva",
			DataNascimento:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Sexo:              "M",
			CategoriaCNH:      &categoriaCNH,
			PaisNascimento:    "BRA",
			PaisNacionalidade: "BRA",
		}
		_, err := u.CreateTrabalhador(ctx, req)
		if err == nil {
			t.Fatal("expected error for invalid categoria_cnh")
		}
	})

	t.Run("valid categoria_cnh", func(t *testing.T) {
		categoriaCNH := "AB"
		req := &CreateTrabalhadorRequest{
			EmpresaID:         "e1",
			CPF:               "11144477735",
			Nome:              "João Silva",
			DataNascimento:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Sexo:              "M",
			CategoriaCNH:      &categoriaCNH,
			PaisNascimento:    "BRA",
			PaisNacionalidade: "BRA",
		}
		resp, err := u.CreateTrabalhador(ctx, req)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if resp == nil {
			t.Fatal("expected response, got nil")
		}
	})

	t.Run("create service error", func(t *testing.T) {
		mockSvc.createErr = errors.New("service error")
		req := &CreateTrabalhadorRequest{
			EmpresaID:         "e1",
			CPF:               "11144477735",
			Nome:              "João Silva",
			DataNascimento:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Sexo:              "M",
			PaisNascimento:    "BRA",
			PaisNacionalidade: "BRA",
		}
		_, err := u.CreateTrabalhador(ctx, req)
		if err == nil {
			t.Fatal("expected error from service")
		}
		mockSvc.createErr = nil
	})
}

// TestUsecaseGetTrabalhador tests get operations
func TestUsecaseGetTrabalhador(t *testing.T) {
	ctx := context.Background()
	u, _ := setupUsecase()

	t.Run("get trabalhador", func(t *testing.T) {
		resp, err := u.GetTrabalhador(ctx, "t1")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if resp == nil {
			t.Fatal("expected response, got nil")
		}
	})
}

// TestUsecaseUpdateTrabalhador tests update operations
func TestUsecaseUpdateTrabalhador(t *testing.T) {
	ctx := context.Background()
	u, _ := setupUsecase()

	t.Run("update without cpf or nis validation", func(t *testing.T) {
		req := &UpdateTrabalhadorRequest{}
		resp, err := u.UpdateTrabalhador(ctx, "t1", req)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if resp == nil {
			t.Fatal("expected response, got nil")
		}
	})

	t.Run("invalid nis in update", func(t *testing.T) {
		nis := "123"
		req := &UpdateTrabalhadorRequest{NIS: &nis}
		_, err := u.UpdateTrabalhador(ctx, "t1", req)
		if err == nil {
			t.Fatal("expected error for invalid NIS")
		}
	})
}

// TestUsecaseDeleteTrabalhador tests delete operations
func TestUsecaseDeleteTrabalhador(t *testing.T) {
	ctx := context.Background()
	u, _ := setupUsecase()

	t.Run("delete trabalhador", func(t *testing.T) {
		err := u.DeleteTrabalhador(ctx, "t1")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	})
}

// TestUsecaseListTrabalhadores tests list operations
func TestUsecaseListTrabalhadores(t *testing.T) {
	ctx := context.Background()
	u, _ := setupUsecase()

	t.Run("list trabalhadores", func(t *testing.T) {
		req := &ListRequest{Page: 1, Limit: 10}
		resp, err := u.ListTrabalhadores(ctx, req)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if resp == nil {
			t.Fatal("expected response, got nil")
		}
	})
}

// TestUsecaseValidateCPF tests CPF validation
func TestUsecaseValidateCPF(t *testing.T) {
	ctx := context.Background()
	u, _ := setupUsecase()

	tests := []struct {
		name    string
		cpf     string
		wantErr bool
	}{
		{"invalid length", "123", true},
		{"all same digits", "11111111111", true},
		{"valid cpf", "11144477735", false},
		{"invalid first check digit", "11244477735", true},
		{"invalid second check digit", "11144477734", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := u.ValidateCPF(ctx, tt.cpf)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateCPF() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestUsecaseValidateTrabalhadorData tests full worker data validation
func TestUsecaseValidateTrabalhadorData(t *testing.T) {
	ctx := context.Background()
	u, _ := setupUsecase()

	t.Run("valid complete data", func(t *testing.T) {
		req := &CreateTrabalhadorRequest{
			EmpresaID:         "e1",
			CPF:               "11144477735",
			Nome:              "João Silva",
			DataNascimento:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Sexo:              "M",
			PaisNascimento:    "BRA",
			PaisNacionalidade: "BRA",
		}
		err := u.ValidateTrabalhadorData(ctx, req)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	})

	t.Run("missing pais_nascimento", func(t *testing.T) {
		req := &CreateTrabalhadorRequest{
			EmpresaID:         "e1",
			CPF:               "11144477735",
			Nome:              "João Silva",
			DataNascimento:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Sexo:              "M",
			PaisNacionalidade: "BRA",
		}
		err := u.ValidateTrabalhadorData(ctx, req)
		if err == nil {
			t.Fatal("expected error for missing pais_nascimento")
		}
	})

	t.Run("missing pais_nacionalidade", func(t *testing.T) {
		req := &CreateTrabalhadorRequest{
			EmpresaID:      "e1",
			CPF:            "11144477735",
			Nome:           "João Silva",
			DataNascimento: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Sexo:           "M",
			PaisNascimento: "BRA",
		}
		err := u.ValidateTrabalhadorData(ctx, req)
		if err == nil {
			t.Fatal("expected error for missing pais_nacionalidade")
		}
	})

	t.Run("invalid estado_civil", func(t *testing.T) {
		estadoCivil := "9"
		req := &CreateTrabalhadorRequest{
			EmpresaID:         "e1",
			CPF:               "11144477735",
			Nome:              "João Silva",
			DataNascimento:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Sexo:              "M",
			EstadoCivil:       &estadoCivil,
			PaisNascimento:    "BRA",
			PaisNacionalidade: "BRA",
		}
		err := u.ValidateTrabalhadorData(ctx, req)
		if err == nil {
			t.Fatal("expected error for invalid estado_civil")
		}
	})

	t.Run("valid estado_civil", func(t *testing.T) {
		estadoCivil := "2"
		req := &CreateTrabalhadorRequest{
			EmpresaID:         "e1",
			CPF:               "11144477735",
			Nome:              "João Silva",
			DataNascimento:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Sexo:              "M",
			EstadoCivil:       &estadoCivil,
			PaisNascimento:    "BRA",
			PaisNacionalidade: "BRA",
		}
		err := u.ValidateTrabalhadorData(ctx, req)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	})

	t.Run("invalid grau_instrucao", func(t *testing.T) {
		grauInstrucao := "99"
		req := &CreateTrabalhadorRequest{
			EmpresaID:         "e1",
			CPF:               "11144477735",
			Nome:              "João Silva",
			DataNascimento:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Sexo:              "M",
			GrauInstrucao:     &grauInstrucao,
			PaisNascimento:    "BRA",
			PaisNacionalidade: "BRA",
		}
		err := u.ValidateTrabalhadorData(ctx, req)
		if err == nil {
			t.Fatal("expected error for invalid grau_instrucao")
		}
	})

	t.Run("valid grau_instrucao", func(t *testing.T) {
		grauInstrucao := "08"
		req := &CreateTrabalhadorRequest{
			EmpresaID:         "e1",
			CPF:               "11144477735",
			Nome:              "João Silva",
			DataNascimento:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Sexo:              "M",
			GrauInstrucao:     &grauInstrucao,
			PaisNascimento:    "BRA",
			PaisNacionalidade: "BRA",
		}
		err := u.ValidateTrabalhadorData(ctx, req)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	})

	t.Run("invalid tipo_deficiencia", func(t *testing.T) {
		tipoDeficiencia := "99"
		req := &CreateTrabalhadorRequest{
			EmpresaID:         "e1",
			CPF:               "11144477735",
			Nome:              "João Silva",
			DataNascimento:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Sexo:              "M",
			TipoDeficiencia:   &tipoDeficiencia,
			PaisNascimento:    "BRA",
			PaisNacionalidade: "BRA",
		}
		err := u.ValidateTrabalhadorData(ctx, req)
		if err == nil {
			t.Fatal("expected error for invalid tipo_deficiencia")
		}
	})

	t.Run("valid tipo_deficiencia", func(t *testing.T) {
		tipoDeficiencia := "3"
		req := &CreateTrabalhadorRequest{
			EmpresaID:         "e1",
			CPF:               "11144477735",
			Nome:              "João Silva",
			DataNascimento:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Sexo:              "M",
			TipoDeficiencia:   &tipoDeficiencia,
			PaisNascimento:    "BRA",
			PaisNacionalidade: "BRA",
		}
		err := u.ValidateTrabalhadorData(ctx, req)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	})

	t.Run("invalid tipo_conta", func(t *testing.T) {
		tipoConta := "5"
		req := &CreateTrabalhadorRequest{
			EmpresaID:         "e1",
			CPF:               "11144477735",
			Nome:              "João Silva",
			DataNascimento:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Sexo:              "M",
			TipoConta:         &tipoConta,
			PaisNascimento:    "BRA",
			PaisNacionalidade: "BRA",
		}
		err := u.ValidateTrabalhadorData(ctx, req)
		if err == nil {
			t.Fatal("expected error for invalid tipo_conta")
		}
	})

	t.Run("valid tipo_conta", func(t *testing.T) {
		tipoConta := "1"
		req := &CreateTrabalhadorRequest{
			EmpresaID:         "e1",
			CPF:               "11144477735",
			Nome:              "João Silva",
			DataNascimento:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Sexo:              "M",
			TipoConta:         &tipoConta,
			PaisNascimento:    "BRA",
			PaisNacionalidade: "BRA",
		}
		err := u.ValidateTrabalhadorData(ctx, req)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	})
}

// TestUsecaseServiceErrors tests service error handling
func TestUsecaseServiceErrors(t *testing.T) {
	ctx := context.Background()
	u, mockSvc2 := setupUsecase()

	t.Run("create service error", func(t *testing.T) {
		mockSvc2.createErr = errors.New("service error")
		req := &CreateTrabalhadorRequest{
			EmpresaID:         "e1",
			CPF:               "11144477735",
			Nome:              "João Silva",
			DataNascimento:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Sexo:              "M",
			PaisNascimento:    "BRA",
			PaisNacionalidade: "BRA",
		}
		_, err := u.CreateTrabalhador(ctx, req)
		if err == nil {
			t.Fatal("expected error from service")
		}
		mockSvc2.createErr = nil
	})

	t.Run("get service error", func(t *testing.T) {
		mockSvc2.getErr = errors.New("service error")
		_, err := u.GetTrabalhador(ctx, "t1")
		if err == nil {
			t.Fatal("expected error from service")
		}
		mockSvc2.getErr = nil
	})

	t.Run("list service error", func(t *testing.T) {
		mockSvc2.listErr = errors.New("service error")
		_, err := u.ListTrabalhadores(ctx, &ListRequest{Page: 1, Limit: 10})
		if err == nil {
			t.Fatal("expected error from service")
		}
		mockSvc2.listErr = nil
	})
}
