package billing_service

import "encoding/json"

type Invoice struct {
	AddigyInvoiceId string      `json:"addigy_invoice_id"`
	Balance         json.Number `json:"balance"`
	Date            string      `json:"date"`
	InvoiceId       string      `json:"invoice_id"`
	InvoiceLink     string      `json:"invoice_link"`
	Provider        string      `json:"provider"`
	Status          string      `json:"status"`
}
