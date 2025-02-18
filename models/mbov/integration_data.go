package mbov

type IntegrationData struct {
	BillingCity    string `json:"billing_city"`
	BillingCountry string `json:"billing_country"`
	BillingState   string `json:"billing_state"`
	BillingStreet  string `json:"billing_street"`
	BillingZipCode string `json:"billing_zip_code"`
	Email          string `json:"email"`
}
