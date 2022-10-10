package main

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/voluntas/webapp"
)

func main() {
	db, err := sql.Open("pgx", webapp.SQLITE_PATH)
	if err != nil {
		log.Fatal("")
	}
	driver, err := pgx.WithInstance(db, &pgx.Config{})
	if err != nil {
		log.Fatal("")
	}

	m, err := migrate.NewWithDatabaseInstance(
		// フォルダが固定なので固定でイイ
		"file://./db/schema",
		"pgx", driver)
	if err != nil {
		log.Fatal("")
	}
	m.Up()
}
