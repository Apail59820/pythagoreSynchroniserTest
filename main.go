package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"pythagoreSynchroniser/config"
	"pythagoreSynchroniser/db"
	"pythagoreSynchroniser/logging"
	"pythagoreSynchroniser/metrics"
	"pythagoreSynchroniser/services"
)

func main() {
	config.Load()
	logging.Setup()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	sqlDB, err := db.ConnectSQL(ctx)
	if err != nil {
		logging.Fatalln("connexion SQL:", err)
	}
	defer sqlDB.Close()

	go func() {
		http.HandleFunc("/", metrics.DashboardHandler(sqlDB))
		logging.Infof("Dashboard disponible sur :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			logging.Errorf("serveur HTTP: %v", err)
		}
	}()

	conn, err := db.Connect(ctx)
	if err != nil {
		logging.Fatalln("Erreur de connexion à la base de données:", err)
	}
	defer func() {
		if err := conn.Close(ctx); err != nil {
			logging.Errorf("fermeture de la connexion : %v", err)
		}
	}()

	logging.Infof("Connexion PostgreSQL établie.")
	logging.Infof("Démarrage du synchroniseur...")

	ticker := time.NewTicker(config.SyncInterval())
	defer ticker.Stop()

	lastID := config.LoadLastID()

	for {
		select {
		case <-ctx.Done():
			logging.Warnf("Arrêt demandé, fermeture...")
			return
		case <-ticker.C:
			start := time.Now()
			invoices, err := db.FetchInvoicesAfterID(ctx, conn, lastID)
			if len(invoices) == 0 {
				logging.Debugf("Aucune nouvelle facture.")
				continue
			}
			if err != nil {
				logging.Errorf("erreur récupération factures: %v", err)
				continue
			}

			for _, inv := range invoices {
				req, err := services.ConvertInvoice(inv)
				if err != nil {
					logging.Errorf("conversion facture %d: %v", inv.ID, err)
					continue
				}
				ref, token, err := services.SendInvoiceToFNE(req, "")
				if err != nil {
					logging.Errorf("envoi FNE facture %d: %v", inv.ID, err)
					//continue
				}
				if err := config.AppendMetadata(config.InvoiceMetadata{
					InvoiceID: inv.ID,
					Reference: ref,
					Token:     token,
				}); err != nil {
					logging.Errorf("sauvegarde metadata facture %d: %v", inv.ID, err)
				}
				if inv.ID > lastID {
					lastID = inv.ID
				}
			}

			if err := config.SaveLastID(lastID); err != nil {
				logging.Errorf("sauvegarde etat: %v", err)
			}
			metrics.RecordSync(time.Since(start))
		}
	}
}
