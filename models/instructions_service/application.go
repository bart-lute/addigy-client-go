package instructions_service

type Application struct {
	Bundle         string `json:"bundle"`
	Custom         bool   `json:"custom"`
	Identifier     string `json:"identifier"`
	IdentifierType string `json:"identifier_type"`
	Name           string `json:"name"`
	Path           string `json:"path"`
	Requirements   string `json:"requirements"`
	Signature      string `json:"signature"`
}
