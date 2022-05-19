package db_handler

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var DBClient *sql.DB

func InitClient() {
	db, err := sql.Open("postgres", "postgres://postgres:postgrespw@localhost:55000/hackernews?sslmode=disable")
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	DBClient = db
}

func HandleMigrate() {
	if err := DBClient.Ping(); err != nil {
		log.Fatal(err)
	}

	driver, _ := postgres.WithInstance(DBClient, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance("file://internal/pkg/db/migrations/postgresql", "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
