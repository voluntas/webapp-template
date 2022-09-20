package main

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"github.com/voluntas/webapp"
)

func main() {
	db, err := sql.Open("sqlite3", webapp.SQLITE_PATH)
	if err != nil {
		log.Fatal("")
	}
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatal("")
	}

	m, err := migrate.NewWithDatabaseInstance(
		// フォルダが固定なので固定でイイ
		"file://./db/schema",
		"sqlite3", driver)
	if err != nil {
		log.Fatal("")
	}
	m.Up()
}
