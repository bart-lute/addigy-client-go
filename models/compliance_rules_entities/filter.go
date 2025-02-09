package compliance_rules_entities

type Filter struct {
	ExcludedIds  []string `json:"excluded_ids"`
	Ids          []string `json:"ids"`
	NameContains string   `json:"name_contains"`
}
