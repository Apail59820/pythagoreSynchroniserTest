package services

import (
	"encoding/json"
	"fmt"
	"pythagoreSynchroniser/models"
	"strconv"
)

// ConvertInvoice convertit une facture de la base en requête FNE.
func ConvertInvoice(inv models.Invoice) (models.FNEInvoiceRequest, error) {
	var req models.FNEInvoiceRequest

	req.InvoiceType = inv.InvoiceType
	req.PaymentMethod = inv.PaymentMethod
	req.Template = inv.Template
	req.IsRne = inv.IsRne
	req.Rne = inv.Rne
	req.ClientNcc = inv.ClientNcc
	req.ClientCompanyName = inv.ClientCompanyName
	req.ClientPhone = strconv.FormatInt(inv.ClientPhone, 10)
	req.ClientEmail = inv.ClientEmail
	req.ClientSellerName = inv.ClientSellerName
	req.PointOfSale = inv.PointOfSale
	req.Establishment = inv.Establishment
	req.CommercialMessage = inv.CommercialMessage
	req.Footer = inv.Footer
	if inv.ForeignCurrency != nil {
		req.ForeignCurrency = *inv.ForeignCurrency
	}
	if inv.ForeignCurrencyRate != nil {
		req.ForeignCurrencyRate = *inv.ForeignCurrencyRate
	}
	req.Taxes = inv.Taxes

	if inv.CustomTaxes != nil {
		var ct []models.CustomTax
		if err := json.Unmarshal([]byte(*inv.CustomTaxes), &ct); err == nil {
			req.CustomTaxes = ct
		}
	}

	if err := json.Unmarshal([]byte(inv.Items), &req.Items); err != nil {
		return req, fmt.Errorf("parse items: %w", err)
	}

	return req, nil
}
