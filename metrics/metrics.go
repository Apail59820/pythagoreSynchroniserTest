package metrics

import (
	"database/sql"
	"html/template"
	"net/http"
	"pythagoreSynchroniser/config"
	"time"
)

type FneMetrics struct {
	TotalInvoices   int
	SentInvoices    int
	ErrorInvoices   int
	ByTemplate      map[string]int
	ByPaymentMethod map[string]int
	ByPointOfSale   map[string]int
	SuccessRate     float64
	AvgSendTime     time.Duration
	LastInvoiceID   int
}

// CollectFneMetrics calcule les statistiques d'envoi FNE.
func CollectFneMetrics(db *sql.DB) (FneMetrics, error) {
	var m FneMetrics
	// total invoices
	if err := db.QueryRow("SELECT COUNT(*) FROM invoices").Scan(&m.TotalInvoices); err != nil {
		return m, err
	}

	// group by template
	rows, err := db.Query("SELECT template, COUNT(*) FROM invoices GROUP BY template")
	if err != nil {
		return m, err
	}
	m.ByTemplate = map[string]int{}
	for rows.Next() {
		var name string
		var count int
		if err := rows.Scan(&name, &count); err != nil {
			rows.Close()
			return m, err
		}
		m.ByTemplate[name] = count
	}
	rows.Close()

	// group by payment method
	rows, err = db.Query("SELECT payment_method, COUNT(*) FROM invoices GROUP BY payment_method")
	if err != nil {
		return m, err
	}
	m.ByPaymentMethod = map[string]int{}
	for rows.Next() {
		var name string
		var count int
		if err := rows.Scan(&name, &count); err != nil {
			rows.Close()
			return m, err
		}
		m.ByPaymentMethod[name] = count
	}
	rows.Close()

	// group by point of sale
	rows, err = db.Query("SELECT point_of_sale, COUNT(*) FROM invoices GROUP BY point_of_sale")
	if err != nil {
		return m, err
	}
	m.ByPointOfSale = map[string]int{}
	for rows.Next() {
		var name string
		var count int
		if err := rows.Scan(&name, &count); err != nil {
			rows.Close()
			return m, err
		}
		m.ByPointOfSale[name] = count
	}
	rows.Close()

	meta, err := config.LoadMetadata()
	if err != nil {
		return m, err
	}
	m.SentInvoices = len(meta)
	m.LastInvoiceID = config.LoadLastID()
	m.ErrorInvoices = m.TotalInvoices - m.SentInvoices
	if m.TotalInvoices > 0 {
		m.SuccessRate = float64(m.SentInvoices) / float64(m.TotalInvoices) * 100
	}
	// AvgSendTime non disponible faute de logs
	return m, nil
}

var dashboardTmpl = template.Must(template.ParseFiles("templates/metrics.html"))

// DashboardHandler renvoie un handler HTTP affichant les métriques.
func DashboardHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m, err := CollectFneMetrics(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := dashboardTmpl.Execute(w, m); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
