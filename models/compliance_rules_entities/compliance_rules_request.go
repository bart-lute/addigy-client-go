package compliance_rules_entities

type ComplianceRulesRequest struct {
	Page          int    `json:"page"`
	PerPage       int    `json:"per_page"`
	Query         Filter `json:"query"`
	SortDirection string
	SortField     string
}
