package empregador_test

import (
	"context"
	"sped/esocial/internal/model/empregador"
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/jmoiron/sqlx"
)

func TestGetEmpresas(t *testing.T) {

	db := sqlx.MustConnect("postgres", "user=postgres password=postgres dbname=esocial sslmode=disable")
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	empresas, err := empregador.GetAll(ctx, db)
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

	fake := faker.New()
	empresa := &empregador.Empresas{
		Id:           "Não pode aparecer no banco",
		Cnpj:         fake.Numerify("04126###0001##"),
		RazaoSocial:  "Empresa Teste sem ID",
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

		err = empregador.InsertOne(ctx, tx, empresa)
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

	for range 5 {
		empresa := empregador.Empresas{
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

		err = empregador.InsertMany(ctx, tx, empresas)
		if err != nil {
			t.Fatalf("Failed to insert empresas: %v", err)
		} else {
			t.Log("Successfully inserted empresas")
		}
	}

}

func TestUpdate(t *testing.T) {

	db := sqlx.MustConnect("postgres", "user=postgres password=postgres dbname=esocial sslmode=disable")
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	empresaToUpdate := empregador.Empresas{
		Id:   "019e14d9-f439-725f-af42-160a214746b3",
		Cnpj: "00000000000001",
	}

	if tx, err := db.BeginTxx(ctx, nil); err != nil {
		t.Fatalf("Failed to begin transaction: %v", err)
	} else {
		defer func() {
			erro := tx.Commit()
			if erro != nil {
				t.Fatalf("Failed to rollback transaction: %v", erro)
			}
		}()

		err = empregador.Update(ctx, tx, &empresaToUpdate)
		if err != nil {
			t.Fatalf("Failed to update empresa: %v", err)
		} else {
			t.Log("Successfully updated empresa")
		}
	}

}

func TestUpdateLote(t *testing.T) {

	db := sqlx.MustConnect("postgres", "user=postgres password=postgres dbname=esocial sslmode=disable")
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	empresasToUpdate := []empregador.Empresas{
		{
			Id:   "019e19a4-c398-718b-8af3-f542bf0d3105",
			Cnpj: "00000000000003",
		},
		{
			Id:   "019e19a4-c398-71bb-b5d9-05887c8a1eb8",
			Cnpj: "00000000000004",
		},
	}

	if tx, err := db.BeginTxx(ctx, nil); err != nil {
		t.Fatalf("Failed to begin transaction: %v", err)
	} else {
		defer func() {
			erro := tx.Commit()
			if erro != nil {
				t.Fatalf("Failed to rollback transaction: %v", erro)
			}
		}()

		err = empregador.UpdateMany(ctx, tx, empresasToUpdate)
		if err != nil {
			t.Fatalf("Failed to update empresas: %v", err)
		} else {
			t.Log("Successfully updated empresas")
		}
	}
}

func TestDeleteLote(t *testing.T) {

	db := sqlx.MustConnect("postgres", "user=postgres password=postgres dbname=esocial sslmode=disable")
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	idsToDelete := []string{
		"019e19a4-c397-7d7e-bb89-ae5847b6e05a",
		"019e19a4-c398-70fa-91eb-187388ce1edb",
		"019e19a4-c398-7157-900d-483fe90a9111",
	}

	if tx, err := db.BeginTxx(ctx, nil); err != nil {
		t.Fatalf("Failed to begin transaction: %v", err)
	} else {
		defer func() {
			erro := tx.Commit()
			if erro != nil {
				t.Fatalf("Failed to rollback transaction: %v", erro)
			}
		}()

		rows, err := empregador.DeleteMany(ctx, tx, idsToDelete)
		if err != nil {
			t.Fatalf("Failed to delete empresas: %v", err)
		} else {
			t.Logf("Successfully deleted empresas - Rows affected: %d", rows)
		}
	}
}

func TestDeleteEmpresa(t *testing.T) {

	db := sqlx.MustConnect("postgres", "user=postgres password=postgres dbname=esocial sslmode=disable")
	defer db.Close()

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

		rows, err := empregador.Delete(ctx, tx, "019e14d9-f439-7253-b529-282d8f5b522d")
		if err != nil {
			t.Fatalf("Failed to delete empresa: %v", err)
		} else {
			t.Logf("Successfully deleted empresa - Rows affected: %d", rows)
		}
	}
}
