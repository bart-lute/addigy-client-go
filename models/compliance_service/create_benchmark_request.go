package compliance_service

type CreateBenchmarkRequest struct {
	ComplianceRulesIds []string `json:"compliance_rules_ids"`
	MaximumOsVersion   string   `json:"maximum_os_version"`
	MinimumOsVersion   string   `json:"minimum_os_version"`
	Name               string   `json:"name"`
	TargetOs           string   `json:"target_os"`
}
