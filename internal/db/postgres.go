package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectPostgres() (*sql.DB, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return nil, fmt.Errorf("la variable de entorno DATABASE_URL no está definida")
	}

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("error al abrir la conexión con PostgreSQL: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error al verificar la conexión con PostgreSQL: %w", err)
	}

	return db, nil
}