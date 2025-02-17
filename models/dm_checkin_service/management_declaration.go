package dm_checkin_service

type ManagementDeclaration struct {
	Active          bool     `json:"active"`
	DeclarationType string   `json:"declaration_type"`
	Identifier      string   `json:"identifier"`
	Reasons         []Reason `json:"reasons"`
	ServerToken     string   `json:"server-token"`
	Valid           string   `json:"valid"`
}
