package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"pythagoreSynchroniser/config"
	"pythagoreSynchroniser/db"
	"pythagoreSynchroniser/metrics"
	"pythagoreSynchroniser/services"
	"time"
)

func main() {
	config.Load()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	sqlDB, err := db.ConnectSQL(ctx)
	if err != nil {
		log.Fatalf("connexion SQL: %v", err)
	}
	defer sqlDB.Close()

	go func() {
		http.HandleFunc("/", metrics.DashboardHandler(sqlDB))
		log.Println("Dashboard disponible sur :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Printf("serveur HTTP: %v", err)
		}
	}()

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

	lastID := config.LoadLastID()

	for {
		select {
		case <-ctx.Done():
			log.Println("Arrêt demandé, fermeture...")
			return
		case <-ticker.C:
			invoices, err := db.FetchInvoicesAfterID(ctx, conn, lastID)
			if len(invoices) == 0 {
				log.Printf("Aucune nouvelle facture.")
				continue
			}
			if err != nil {
				log.Printf("erreur récupération factures: %v", err)
				continue
			}

			for _, inv := range invoices {
				req, err := services.ConvertInvoice(inv)
				if err != nil {
					log.Printf("conversion facture %d: %v", inv.ID, err)
					continue
				}
				ref, token, err := services.SendInvoiceToFNE(req, "")
				if err != nil {
					log.Printf("envoi FNE facture %d: %v", inv.ID, err)
					//continue
				}
				if err := config.AppendMetadata(config.InvoiceMetadata{
					InvoiceID: inv.ID,
					Reference: ref,
					Token:     token,
				}); err != nil {
					log.Printf("sauvegarde metadata facture %d: %v", inv.ID, err)
				}
				if inv.ID > lastID {
					lastID = inv.ID
				}
			}

			if err := config.SaveLastID(lastID); err != nil {
				log.Printf("sauvegarde etat: %v", err)
			}
		}
	}
}
