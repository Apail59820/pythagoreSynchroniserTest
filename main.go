package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"pythagoreSynchroniser/config"
	"pythagoreSynchroniser/db"
	"time"
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
			start, err := config.StartDate()
			if err != nil {
				log.Printf("date de début: %v", err)
				continue
			}
			end, err := config.EndDate()
			if err != nil {
				log.Printf("date de fin: %v", err)
				continue
			}

			invoices, err := db.FetchInvoicesBetween(ctx, conn, start, end)
			if err != nil {
				log.Printf("erreur récupération factures: %v", err)
				continue
			}

			for _, inv := range invoices {
				b, err := json.Marshal(inv)
				if err != nil {
					log.Printf("marshal facture %d: %v", inv.ID, err)
					continue
				}
				log.Println(string(b))
			}
		}
	}
}
