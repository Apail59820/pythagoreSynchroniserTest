package main

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"pythagoreSynchroniser/config"
	"pythagoreSynchroniser/db"
	"time"
)

func main() {
	config.LoadEnv()

	conn, err := db.ConnectToPostgres()
	ctx := context.Background()

	if err != nil {
		log.Fatalf("Erreur de connexion à la base de données : %v", err)
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {

		}
	}(conn, ctx)

	log.Println("Connexion PostgreSQL établie.")
	log.Println("Démarrage du synchroniseur...")

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Println("Récupération des factures")
		}
	}
}
