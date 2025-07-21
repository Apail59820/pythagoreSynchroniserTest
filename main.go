package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/jackc/pgx/v5"

	"pythagoreSynchroniser/config"
	"pythagoreSynchroniser/db"
)

func main() {
	config.Load()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	conn, err := db.Connect(ctx)
	if err != nil {
		log.Fatalf("Erreur de connexion à la base de données : %v", err)
	}
	defer func() {
		if err := conn.Close(ctx); err != nil {
			log.Printf("fermeture de la connexion : %v", err)
		}
	}()

	log.Println("Connexion PostgreSQL établie.")
	log.Println("Démarrage du synchroniseur...")

	ticker := time.NewTicker(config.SyncInterval())
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("Arrêt demandé, fermeture...")
			return
		case <-ticker.C:
			log.Println("Récupération des factures")
		}
	}
}
