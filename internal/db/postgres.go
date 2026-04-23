package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectPostgres() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=postgres password=MartinPSQL dbname=series_tracker_db sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error al abrir la conexión con PostgreSQL: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error al verificar la conexión con PostgreSQL: %w", err)
	}

	return db, nil
}