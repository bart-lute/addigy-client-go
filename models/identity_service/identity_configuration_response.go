package identity_service

type IdentityConfigurationResponse struct {
	Count   int                   `json:"count"`
	Results IdentityConfiguration `json:"results"`
	Total   int                   `json:"total"`
}
