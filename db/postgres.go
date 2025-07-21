package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

// Connect ouvre une connexion PostgreSQL en utilisant les variables d'environnement.
func Connect(ctx context.Context) (*pgx.Conn, error) {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
