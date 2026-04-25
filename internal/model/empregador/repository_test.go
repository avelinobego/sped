package empregador_test

import (
	"sped/esocial/internal/model/empregador"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jaswdr/faker"
)

func TestGetEmpresas(t *testing.T) {

	repo := empregador.NewRepositoryPg()
	err := repo.Connect()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer repo.Close()

	empresas, err := repo.GetEmpresas()
	if err != nil {
		t.Fatalf("Failed to get empresas: %v", err)
	}
	if len(empresas) == 0 {
		t.Fatal("Expected at least one empresa, got zero")
	} else {
		t.Logf("Successfully retrieved %d empresas", len(empresas))
		for _, empresa := range empresas {
			t.Logf("%+v", empresa)
		}
	}
}

func TestInsert(t *testing.T) {
	repo := empregador.NewRepositoryPg()
	err := repo.Connect()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer repo.Close()

	uuid, err := uuid.NewV7()
	if err != nil {
		t.Fatalf("Failed to generate UUID: %v", err)
	}

	fake := faker.New()
	empresa := &empregador.Empresas{
		Id:           uuid.String(),
		Cnpj:         fake.Numerify("04126###0001##"),
		RazaoSocial:  "Empresa Teste",
		DataCadastro: new(time.Now()),
	}

	if err = repo.BeginTx(); err != nil {
		t.Fatalf("Failed to begin transaction: %v", err)
	}
	defer repo.RollbackTx()
	err = repo.Insert(empresa)
	if err != nil {
		t.Fatalf("Failed to insert empresa: %v", err)
	} else {
		t.Log("Successfully inserted empresa")
	}
}
