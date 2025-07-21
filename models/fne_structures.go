package models

type FNEInvoiceRequest struct {
	InvoiceType         string      `json:"invoiceType"`
	PaymentMethod       string      `json:"paymentMethod"`
	Template            string      `json:"template"`
	IsRne               bool        `json:"isRne"`
	Rne                 *string     `json:"rne,omitempty"`
	ClientNcc           *string     `json:"clientNcc,omitempty"`
	ClientCompanyName   string      `json:"clientCompanyName"`
	ClientPhone         string      `json:"clientPhone"`
	ClientEmail         string      `json:"clientEmail"`
	ClientSellerName    *string     `json:"clientSellerName,omitempty"`
	PointOfSale         string      `json:"pointOfSale"`
	Establishment       string      `json:"establishment"`
	CommercialMessage   *string     `json:"commercialMessage,omitempty"`
	Footer              *string     `json:"footer,omitempty"`
	ForeignCurrency     string      `json:"foreignCurrency,omitempty"`
	ForeignCurrencyRate float64     `json:"foreignCurrencyRate,omitempty"`
	Items               []FNEItem   `json:"items"`
	Taxes               string      `json:"taxes"`
	CustomTaxes         []CustomTax `json:"customTaxes,omitempty"`
	Discount            float64     `json:"discount,omitempty"` // total discount
}

type FNEItem struct {
	Reference       string      `json:"reference"`
	Description     string      `json:"description"`
	Quantity        float64     `json:"quantity"`
	Amount          float64     `json:"amount"`
	Discount        float64     `json:"discount,omitempty"`
	MeasurementUnit string      `json:"measurementUnit,omitempty"`
	Taxes           []string    `json:"taxes"`
	CustomTaxes     []CustomTax `json:"customTaxes,omitempty"`
}

type CustomTax struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}
