package empregador_test

import (
	"context"
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	empresas, err := empregador.GetEmpresas(ctx, db)
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if tx, err := db.BeginTxx(ctx, nil); err != nil {
		t.Fatalf("Failed to begin transaction: %v", err)
	} else {
		defer func() {
			erro := tx.Commit()
			if erro != nil {
				t.Fatalf("Failed to rollback transaction: %v", erro)
			}
		}()

		err = empregador.InsertEmpresa(ctx, tx, empresa)
		if err != nil {
			t.Fatalf("Failed to insert empresa: %v", err)
		} else {
			t.Log("Successfully inserted empresa")
		}
	}

}

func TestInsertLote(t *testing.T) {

	db := sqlx.MustConnect("postgres", "user=postgres password=postgres dbname=esocial sslmode=disable")
	defer db.Close()

	var empresas []empregador.Empresas
	fake := faker.New()

	for i := 0; i < 5; i++ {
		uuid, err := uuid.NewV7()
		if err != nil {
			t.Fatalf("Failed to generate UUID: %v", err)
		}

		empresa := empregador.Empresas{
			Id:           uuid.String(),
			Cnpj:         fake.Numerify("04126###0001##"),
			RazaoSocial:  "Empresa Teste",
			DataCadastro: new(time.Now()),
		}
		empresas = append(empresas, empresa)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if tx, err := db.BeginTxx(ctx, nil); err != nil {
		t.Fatalf("Failed to begin transaction: %v", err)
	} else {
		defer func() {
			erro := tx.Commit()
			if erro != nil {
				t.Fatalf("Failed to rollback transaction: %v", erro)
			}
		}()

		err = empregador.InsertEmpresas(ctx, tx, empresas)
		if err != nil {
			t.Fatalf("Failed to insert empresas: %v", err)
		} else {
			t.Log("Successfully inserted empresas")
		}
	}

}
