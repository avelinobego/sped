package empregador

import (
	"fmt"
	"sped/esocial/pkg/utils"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type reposiroyPg struct {
	db *sqlx.DB
	tx *sqlx.Tx
}

func NewRepositoryPg() *reposiroyPg {
	return &reposiroyPg{}
}

func (c *reposiroyPg) Connect() error {
	var err error
	//TODO: tornar isso genérico
	c.db, err = sqlx.Connect("postgres", "user=postgres password=postgres dbname=esocial sslmode=disable")
	return err
}

func (c *reposiroyPg) BeginTx() error {
	var err error
	c.tx, err = c.db.Beginx()
	return err
}

func (c *reposiroyPg) CommitTx() error {
	if c.tx == nil {
		return fmt.Errorf("transaction not open")
	}
	return c.tx.Commit()
}

func (c *reposiroyPg) RollbackTx() error {
	if c.tx == nil {
		return fmt.Errorf("transaction not open")
	}
	return c.tx.Rollback()
}

func (c *reposiroyPg) Close() error {
	return c.db.Close()
}

func (c *reposiroyPg) GetEmpresas() ([]Empresas, error) {
	var empresas []Empresas
	err := c.db.Select(&empresas, "SELECT * FROM empresas")
	if err != nil {
		return nil, err
	}
	return empresas, nil
}

func (c *reposiroyPg) Insert(empresa *Empresas) error {
	query, err := utils.BuildInsert("empresas", empresa)
	if err != nil {
		return err
	}
	_, err = c.db.NamedExec(query, empresa)
	return err
}
