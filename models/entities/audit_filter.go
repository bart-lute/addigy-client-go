package entities

import "encoding/json"

type AuditFilter struct {
	AuditField   string      `json:"audit_field"`
	BooleanValue bool        `json:"boolean_value"`
	DateValue    string      `json:"date_value"`
	ListValue    []string    `json:"list_value"`
	NumberValue  json.Number `json:"number_value"`
	Operation    string      `json:"operation"`
	StringValue  string      `json:"string_value"`
	Type         string      `json:"type"`
}
