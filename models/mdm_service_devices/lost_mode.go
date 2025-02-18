package mdm_service_devices

type LostMode struct {
	Footnote     string       `json:"footnote"`
	LastDisabled string       `json:"last_disabled"`
	LastEnabled  string       `json:"last_enabled"`
	LastLocation LocationInfo `json:"last_location"`
	Message      string       `json:"message"`
	PhoneNumber  string       `json:"phone_number"`
}
