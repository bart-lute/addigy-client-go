package billing_service

type Subscription struct {
	BillingFrequency string `json:"billing_frequency"`
	Name             string `json:"name"`
	NextPaymentDate  string `json:"next_payment_date"`
	PaymentMethod    string `json:"payment_method"`
	//Plans	[billing_service.Plan{...}]
	StartDate string `json:"start_date"`
	Status    string `json:"status"`
}
