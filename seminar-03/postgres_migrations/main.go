package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func main() {
	dsn := "postgres://postgres:postgres@localhost:5432/app?sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatal(err)
	}

	log.Println("Migrations applied")
}
