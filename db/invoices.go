package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"pythagoreSynchroniser/models"
)

// FetchInvoicesBetween retourne toutes les factures dont la date de creation
// est comprise entre start et end.
func FetchInvoicesBetween(ctx context.Context, conn *pgx.Conn, start, end time.Time) ([]models.Invoice, error) {
	const query = `
        SELECT id, invoice_type, payment_method, template, is_rne, rne,
               client_ncc, client_company_name, client_phone, client_email,
               client_seller_name, point_of_sale, establishment,
               commercial_message, footer, foreign_currency,
               foreign_currency_rate, taxes, custom_taxes, items
        FROM invoices
        WHERE created_at >= $1 AND created_at <= $2
    `

	rows, err := conn.Query(ctx, query, start, end)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	var invoices []models.Invoice
	for rows.Next() {
		var inv models.Invoice
		if err := rows.Scan(
			&inv.ID,
			&inv.InvoiceType,
			&inv.PaymentMethod,
			&inv.Template,
			&inv.IsRne,
			&inv.Rne,
			&inv.ClientNcc,
			&inv.ClientCompanyName,
			&inv.ClientPhone,
			&inv.ClientEmail,
			&inv.ClientSellerName,
			&inv.PointOfSale,
			&inv.Establishment,
			&inv.CommercialMessage,
			&inv.Footer,
			&inv.ForeignCurrency,
			&inv.ForeignCurrencyRate,
			&inv.Taxes,
			&inv.CustomTaxes,
			&inv.Items,
		); err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}
		invoices = append(invoices, inv)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return invoices, nil
}

// FetchInvoicesAfterID retourne toutes les factures avec un identifiant superieur a lastID.
func FetchInvoicesAfterID(ctx context.Context, conn *pgx.Conn, lastID int) ([]models.Invoice, error) {
	const query = `
        SELECT id, invoice_type, payment_method, template, is_rne, rne,
               client_ncc, client_company_name, client_phone, client_email,
               client_seller_name, point_of_sale, establishment,
               commercial_message, footer, foreign_currency,
               foreign_currency_rate, taxes, custom_taxes, items
        FROM invoices
        WHERE id > $1
        ORDER BY id 
    `

	rows, err := conn.Query(ctx, query, lastID)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	var invoices []models.Invoice
	for rows.Next() {
		var inv models.Invoice
		if err := rows.Scan(
			&inv.ID,
			&inv.InvoiceType,
			&inv.PaymentMethod,
			&inv.Template,
			&inv.IsRne,
			&inv.Rne,
			&inv.ClientNcc,
			&inv.ClientCompanyName,
			&inv.ClientPhone,
			&inv.ClientEmail,
			&inv.ClientSellerName,
			&inv.PointOfSale,
			&inv.Establishment,
			&inv.CommercialMessage,
			&inv.Footer,
			&inv.ForeignCurrency,
			&inv.ForeignCurrencyRate,
			&inv.Taxes,
			&inv.CustomTaxes,
			&inv.Items,
		); err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}
		invoices = append(invoices, inv)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return invoices, nil
}
