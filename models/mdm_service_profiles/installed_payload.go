package mdm_service_profiles

import "encoding/json"

type InstalledPayload struct {
	Description  string      `json:"description"`
	DisplayName  string      `json:"display_name"`
	Identifier   string      `json:"identifier"`
	Organization string      `json:"organization"`
	Priority     json.Number `json:"priority"`
	Type         string      `json:"type"`
	Uuid         string      `json:"uuid"`
	Version      interface{} `json:"version"`
}
