package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

var ctx = context.Background()

func ConnectToPostgres() (*pgx.Conn, error) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
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
