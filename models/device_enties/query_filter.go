package device_enties

type QueryFilter struct {
	Filters   []AuditFilter `json:"filters,omitempty"`
	PolicyId  string        `json:"policyId,omitempty"`
	SearchAny string        `json:"search_any,omitempty"`
}
