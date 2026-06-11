package connection

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/lib/pq" // To register the driver.
)

var (
	result *sql.DB
	once   sync.Once
)

func GetDB() *sql.DB {
	once.Do(func() {
		result = connect()
	})
	return result
}

func connect() *sql.DB {
	db, err := sql.Open("postgres", "host=localhost dbname=esocial connect_timeout=5")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}
