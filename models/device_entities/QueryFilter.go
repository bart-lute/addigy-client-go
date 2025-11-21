package device_entities

type QueryFilter struct {
	Filters   *[]AuditFilter `json:"filters"`
	PolicyId  string         `json:"policy_id"`
	SearchAny string         `json:"search_any"`
}
