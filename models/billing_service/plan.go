package billing_service

type Plan struct {
	Name     string `json:"name"`
	Price    string `json:"price"`
	Quantity int    `json:"quantity"`
}
