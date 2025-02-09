package billing_service

type BillingAccount struct {
	BillingCity       string `json:"billing_city"`
	BillingCountry    string `json:"billing_country"`
	BillingPostalCode string `json:"billing_postal_code"`
	BillingState      string `json:"billing_state"`
	BillingStreet     string `json:"billing_street"`
	Id                string `json:"id"`
	Name              string `json:"name"`
}
