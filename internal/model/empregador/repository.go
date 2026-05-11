package empregador

import (
	"context"
	"sped/esocial/pkg/utils"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var vazio = Empresas{}

func GetEmpresas(context context.Context, q *sqlx.DB) (result []Empresas, err error) {
	err = q.SelectContext(context, &result, "SELECT * FROM empresas")
	return
}

func InsertEmpresa(context context.Context, tx *sqlx.Tx, empresa *Empresas) error {
	query, err := utils.BuildInsert("empresas", vazio)
	if err != nil {
		return err
	}

	_, err = tx.NamedExecContext(context, query, empresa)
	return err
}

func InsertEmpresas(context context.Context, tx *sqlx.Tx, empresas []Empresas) error {
	query, err := utils.BuildInsert("empresas", vazio)
	if err != nil {
		return err
	}
	_, err = tx.NamedExecContext(context, query, empresas)
	return err
}
