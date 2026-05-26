package repository

import (
	"context"
	"sped/esocial/internal/model/empregador/entity"
	"sped/esocial/pkg/generatesql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var vazio = entity.Empresas{}

func GetAll(context context.Context, q *sqlx.DB) (result []entity.Empresas, err error) {
	err = q.SelectContext(context, &result, "SELECT * FROM empresas")
	return
}

func GetById(context context.Context, q *sqlx.DB, id string) (result entity.Empresas, err error) {
	err = q.GetContext(context, &result, "SELECT * FROM empresas WHERE id = $1", id)
	return
}

func GetByCNPJ(context context.Context, q *sqlx.DB, cnpj string) (result entity.Empresas, err error) {
	err = q.GetContext(context, &result, "SELECT * FROM empresas WHERE cnpj = $1", cnpj)
	return
}

func InsertOne(context context.Context, tx *sqlx.Tx, empresa *entity.Empresas) error {
	query, err := generatesql.BuildInsert("empresas", vazio)
	if err != nil {
		return err
	}

	_, err = tx.NamedExecContext(context, query, empresa)
	return err
}

func InsertMany(context context.Context, tx *sqlx.Tx, empresas []entity.Empresas) error {
	query, err := generatesql.BuildInsert("empresas", vazio)
	if err != nil {
		return err
	}
	_, err = tx.NamedExecContext(context, query, empresas)
	return err
}

func Update(context context.Context, tx *sqlx.Tx, empresa *entity.Empresas) error {
	query, err := generatesql.BuildUpdate("empresas", vazio, "id")
	if err != nil {
		return err
	}
	_, err = tx.NamedExecContext(context, query, empresa)
	return err
}

func UpdateMany(context context.Context, tx *sqlx.Tx, empresas []entity.Empresas) error {
	for _, empresa := range empresas {
		if err := Update(context, tx, &empresa); err != nil {
			return err
		}
	}
	return nil
}

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
