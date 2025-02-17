package identity_service

type Services struct {
	Azure  AzureProvider  `json:"azure"`
	Google GoogleProvider `json:"google"`
	Okta   OktaProvider   `json:"okta"`
}
