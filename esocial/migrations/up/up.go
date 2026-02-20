package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("../.migration")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := os.Getenv("DATABASE_URL")
	m, err := migrate.New(
		"file://../scripts",
		url)
	if err != nil {
		log.Fatal(err)
	}

	defer m.Close()

	err = m.Up()
	if err != nil {
		log.Fatal(err)
	}
}
