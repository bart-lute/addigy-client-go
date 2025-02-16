package correlation_engine

type Rule struct {
	Created  string `json:"created"`
	Disabled bool   `json:"disabled"`
	RuleId   string `json:"rule_id "`
	Script   string `json:"script"`
}
