package compliance_service

type UpdateBenchmarkRequest struct {
	ComplianceRulesIds []string `json:"compliance_rules_ids"`
	Id                 string   `json:"id"`
	MaximumOsVersion   string   `json:"maximum_os_version"`
	MinimumOsVersion   string   `json:"minimum_os_version"`
	Name               string   `json:"name"`
	TargetOs           string   `json:"target_os"`
}
