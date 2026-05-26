package services

import (
	"context"
	"sped/esocial/internal/model/empregador/entity"
	repo_empregador "sped/esocial/internal/model/empregador/repository"
	"sped/esocial/pkg/config"
)

// Retorna todas as empresas cadastradas
func All() (result []entity.Empresas, err error) {
	context := context.Background()
	result, err = repo_empregador.GetAll(context, config.GetDB())
	return
}

// Retorna uma empresa pelo id
func ByID(id string) (result entity.Empresas, err error) {
	context := context.Background()
	result, err = repo_empregador.GetById(context, config.GetDB(), id)
	return
}

func ByCNPJ(cnpj string) (result entity.Empresas, err error) {
	context := context.Background()
	result, err = repo_empregador.GetByCNPJ(context, config.GetDB(), cnpj)
	return
}

func Insert(empresa *entity.Empresas) error {
	context := context.Background()
	tx, err := config.GetDB().BeginTxx(context, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := repo_empregador.InsertOne(context, tx, empresa); err != nil {
		return err
	}

	return tx.Commit()
}

func InsertMany(empresas []entity.Empresas) error {
	context := context.Background()
	tx, err := config.GetDB().BeginTxx(context, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := repo_empregador.InsertMany(context, tx, empresas); err != nil {
		return err
	}

	return tx.Commit()
}

func Update(empresa *entity.Empresas) error {
	context := context.Background()
	tx, err := config.GetDB().BeginTxx(context, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := repo_empregador.Update(context, tx, empresa); err != nil {
		return err
	}

	return tx.Commit()
}

func UpdateMany(empresas []entity.Empresas) error {
	context := context.Background()
	tx, err := config.GetDB().BeginTxx(context, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := repo_empregador.UpdateMany(context, tx, empresas); err != nil {
		return err
	}

	return tx.Commit()
}

func DeleteMany(ids []string) (rows int64, err error) {
	context := context.Background()
	tx, err := config.GetDB().BeginTxx(context, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	rows, err = repo_empregador.DeleteMany(context, tx, ids)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	return rows, err
}

func Delete(id string) (rows int64, err error) {
	context := context.Background()
	tx, err := config.GetDB().BeginTxx(context, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	rows, err = repo_empregador.Delete(context, tx, id)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	return rows, err
}
