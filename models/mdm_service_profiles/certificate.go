package mdm_service_profiles

type Certificate struct {
	ExpirationDate string `json:"expiration_date"`
	Id             string `json:"id"`
	IsExpired      bool   `json:"is_expired"`
	Name           string `json:"name"`
	Topic          string `json:"topic"`
}
