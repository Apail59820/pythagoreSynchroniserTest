package models

type Invoice struct {
	ID                  int
	InvoiceType         string
	PaymentMethod       string
	Template            string
	IsRne               bool
	Rne                 *string
	ClientNcc           *string
	ClientCompanyName   string
	ClientPhone         int64
	ClientEmail         string
	ClientSellerName    *string
	PointOfSale         string
	Establishment       string
	CommercialMessage   *string
	Footer              *string
	ForeignCurrency     *string
	ForeignCurrencyRate *float64
	Taxes               string
	CustomTaxes         *string
	Items               string
}
