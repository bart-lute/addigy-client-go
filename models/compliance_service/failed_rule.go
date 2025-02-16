package compliance_service

type FailedRule struct {
	Error                string `json:"error"`
	RemediationCommandId string `json:"remediation_command_id"`
	RuleId               string `json:"rule_id"`
}
