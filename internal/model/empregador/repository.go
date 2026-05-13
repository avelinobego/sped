package empregador

import (
	"context"
	"sped/esocial/pkg/utils"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var vazio = Empresas{}

func GetAll(context context.Context, q *sqlx.DB) (result []Empresas, err error) {
	err = q.SelectContext(context, &result, "SELECT * FROM empresas")
	return
}

func GetEmpresaById(context context.Context, q *sqlx.DB, id string) (result Empresas, err error) {
	err = q.GetContext(context, &result, "SELECT * FROM empresas WHERE id = $1", id)
	return
}

func InsertOne(context context.Context, tx *sqlx.Tx, empresa *Empresas) error {
	query, err := utils.BuildInsert("empresas", vazio)
	if err != nil {
		return err
	}

	_, err = tx.NamedExecContext(context, query, empresa)
	return err
}

func InsertMany(context context.Context, tx *sqlx.Tx, empresas []Empresas) error {
	query, err := utils.BuildInsert("empresas", vazio)
	if err != nil {
		return err
	}
	_, err = tx.NamedExecContext(context, query, empresas)
	return err
}

func Update(context context.Context, tx *sqlx.Tx, empresa *Empresas) error {
	query, err := utils.BuildUpdate("empresas", vazio, "id")
	if err != nil {
		return err
	}
	_, err = tx.NamedExecContext(context, query, empresa)
	return err
}

func UpdateMany(context context.Context, tx *sqlx.Tx, empresas []Empresas) error {
	for _, empresa := range empresas {
		if err := Update(context, tx, &empresa); err != nil {
			return err
		}
	}
	return nil
}

// func PartialUpdate(context context.Context, tx *sqlx.Tx, empresa *Empresas) error {
// 	query, args, err := utils.BuildPartialUpdate("empresas", vazio, "id", empresa)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = tx.ExecContext(context, query, args...)
// 	return err
// }

func DeleteMany(context context.Context, tx *sqlx.Tx, ids []string) (rows int64, err error) {
	query, args, err := sqlx.In("DELETE FROM empresas WHERE id IN (?)", ids)
	if err != nil {
		return
	}
	query = tx.Rebind(query)
	r, err := tx.ExecContext(context, query, args...)
	if err != nil {
		return
	}
	rows, err = r.RowsAffected()
	return
}

func Delete(context context.Context, tx *sqlx.Tx, id string) (rows int64, err error) {
	r, err := tx.ExecContext(context, "DELETE FROM empresas WHERE id = $1", id)
	if err != nil {
		return
	}
	rows, err = r.RowsAffected()
	return
}
