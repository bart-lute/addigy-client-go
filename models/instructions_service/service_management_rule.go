package instructions_service

type ServiceManagementRule struct {
	Comment   string `json:"comment"`
	RuleType  string `json:"rule_type"`
	RuleValue string `json:"rule_value"`
}
