package device_entities

type AuditFilter struct {
	AuditField string `json:"audit_field"`
	Operation  string `json:"operation"`
	RangeValue any    `json:"range_value"`
	Type       string `json:"type"`
	Value      any    `json:"value"`
}
