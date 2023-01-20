package db

import (
	"database/sql"

	_ "embed"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var (
	//go:embed postgres/piglatin_db.sql
	migration string
)

func SetupDatabase() (*sql.DB, error) {
	dsn := "postgres://piglatin:piglatin@localhost:5432/piglatin?sslmode=disable"
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(migration)

	return db, err
}
