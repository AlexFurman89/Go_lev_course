package db

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewDB() (*sql.DB, error) {
	dsn := "postgres://alexfurman@localhost:5432/contacts_db?sslmode=disable"

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("open error: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping error: %w", err)
	}

	return db, nil
}
