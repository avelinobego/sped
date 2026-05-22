package empregador

import (
	"context"
	"sped/esocial/pkg/config"
)

// Retorna todas as empresas cadastradas
func AllEmpresas() (result []Empresas, err error) {
	context := context.Background()
	result, err = GetAll(context, config.GetDB())
	return
}

// Retorna uma empresa pelo id
func GetEmpresaByID(id string) (result Empresas, err error) {
	context := context.Background()
	result, err = GetById(context, config.GetDB(), id)
	return
}

func GetEmpresaByCNPJ(cnpj string) (result Empresas, err error) {
	context := context.Background()
	result, err = GetByCNPJ(context, config.GetDB(), cnpj)
	return
}

func InsertEmpresa(empresa *Empresas) error {
	context := context.Background()
	tx, err := config.GetDB().BeginTxx(context, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := InsertOne(context, tx, empresa); err != nil {
		return err
	}

	return tx.Commit()
}

func InsertEmpresas(empresas []Empresas) error {
	context := context.Background()
	tx, err := config.GetDB().BeginTxx(context, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := InsertMany(context, tx, empresas); err != nil {
		return err
	}

	return tx.Commit()
}

func UpdateEmpresa(empresa *Empresas) error {
	context := context.Background()
	tx, err := config.GetDB().BeginTxx(context, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := Update(context, tx, empresa); err != nil {
		return err
	}

	return tx.Commit()
}

func UpdateEmpresas(empresas []Empresas) error {
	context := context.Background()
	tx, err := config.GetDB().BeginTxx(context, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := UpdateMany(context, tx, empresas); err != nil {
		return err
	}

	return tx.Commit()
}

func DeleteEmpresas(ids []string) (rows int64, err error) {
	context := context.Background()
	tx, err := config.GetDB().BeginTxx(context, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	rows, err = DeleteMany(context, tx, ids)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	return rows, err
}

func DeleteEmpresa(id string) (rows int64, err error) {
	context := context.Background()
	tx, err := config.GetDB().BeginTxx(context, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	rows, err = Delete(context, tx, id)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	return rows, err
}
