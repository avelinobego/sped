package config

import (
	"sync"

	"github.com/jmoiron/sqlx"
)

var (
	db   *sqlx.DB
	once sync.Once
)

func GetDB() *sqlx.DB {
	once.Do(func() {
		//TODO: Move connection string to env variables or config file
		db = sqlx.MustConnect("postgres", "user=postgres password=postgres dbname=esocial sslmode=disable")
	})
	return db
}
