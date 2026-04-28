package empregador_test

import (
	"sped/esocial/internal/model/empregador"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jaswdr/faker"
	"github.com/jmoiron/sqlx"
)

func TestGetEmpresas(t *testing.T) {

	db := sqlx.MustConnect("postgres", "user=postgres password=postgres dbname=esocial sslmode=disable")
	defer db.Close()

	empresas, err := empregador.GetEmpresas(db)
	if err != nil {
		t.Fatalf("Failed to get empresas: %v", err)
	}

	if len(empresas) == 0 {
		t.Fatal("Expected at least one empresa, got zero")
	} else {
		t.Logf("Successfully retrieved %d empresas", len(empresas))
		for _, empresa := range empresas {
			t.Logf("ID: %s, CNPJ: %s, Razao Social: %s", empresa.Id, empresa.Cnpj, empresa.RazaoSocial)
		}
	}
}

func TestInsert(t *testing.T) {

	db := sqlx.MustConnect("postgres", "user=postgres password=postgres dbname=esocial sslmode=disable")
	defer db.Close()

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

	if tx, err := db.Beginx(); err != nil {
		t.Fatalf("Failed to begin transaction: %v", err)
	} else {
		defer func() {
			erro := tx.Commit()
			if erro != nil {
				t.Fatalf("Failed to rollback transaction: %v", erro)
			}
		}()

		err = empregador.InsertEmpresa(tx, empresa)
		if err != nil {
			t.Fatalf("Failed to insert empresa: %v", err)
		} else {
			t.Log("Successfully inserted empresa")
		}
	}

}
