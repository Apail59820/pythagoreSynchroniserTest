package services

import (
	"pythagoreSynchroniser/models"
	"testing"
)

func TestConvertInvoice(t *testing.T) {
	itemsJSON := `[{"reference":"ref1","description":"desc","quantity":1,"amount":10,"taxes":["A"]}]`
	customTaxes := `[{"name":"tva","amount":1}]`
	rate := 2.5
	foreign := "USD"
	seller := "seller"
	inv := models.Invoice{
		ID:                  1,
		InvoiceType:         "sale",
		PaymentMethod:       "cash",
		Template:            "default",
		IsRne:               true,
		Rne:                 nil,
		ClientNcc:           nil,
		ClientCompanyName:   "ACME",
		ClientPhone:         12345,
		ClientEmail:         "a@b.com",
		ClientSellerName:    &seller,
		PointOfSale:         "pos",
		Establishment:       "est",
		CommercialMessage:   nil,
		Footer:              nil,
		ForeignCurrency:     &foreign,
		ForeignCurrencyRate: &rate,
		Taxes:               "VAT",
		CustomTaxes:         &customTaxes,
		Items:               itemsJSON,
	}

	req, err := ConvertInvoice(inv)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if req.InvoiceType != inv.InvoiceType || req.PaymentMethod != inv.PaymentMethod || len(req.Items) != 1 {
		t.Fatalf("unexpected request: %+v", req)
	}
	if req.ForeignCurrency != foreign || req.ForeignCurrencyRate != rate {
		t.Fatalf("currency not copied")
	}
	if len(req.CustomTaxes) != 1 || req.CustomTaxes[0].Name != "tva" {
		t.Fatalf("custom taxes not parsed")
	}
}

func TestConvertInvoiceBadItems(t *testing.T) {
	inv := models.Invoice{Items: "not-json"}
	if _, err := ConvertInvoice(inv); err == nil {
		t.Fatal("expected error")
	}
}
