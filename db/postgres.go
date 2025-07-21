package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
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

// ConnectSQL ouvre une connexion *sql.DB pour les statistiques.
func ConnectSQL(ctx context.Context) (*sql.DB, error) {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	cfg, err := pgx.ParseConfig(dbURL)
	if err != nil {
		return nil, err
	}
	db := stdlib.OpenDB(*cfg)
	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
