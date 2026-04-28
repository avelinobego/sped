package empregador

import (
	"sped/esocial/pkg/utils"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetEmpresas(q *sqlx.DB) (result []Empresas, err error) {
	err = q.Select(&result, "SELECT * FROM empresas")
	return
}

func InsertEmpresa(tx *sqlx.Tx, empresa *Empresas) error {
	query, err := utils.BuildInsert("empresas", empresa)
	if err != nil {
		return err
	}

	_, err = tx.NamedExec(query, empresa)
	return err
}

func InsertEmpresas(tx *sqlx.Tx, empresas []Empresas) error {
	//TODO: Implementar inserção em lote
	return nil
}
