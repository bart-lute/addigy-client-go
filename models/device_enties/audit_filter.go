package device_enties

type AuditFilter struct {
	AuditField string `json:"audit_field,omitempty"`
	Operation  string `json:"operation,omitempty"`
	RangeValue string `json:"range_value,omitempty"`
	Type       string `json:"type,omitempty"`
	Value      string `json:"value,omitempty"`
}
