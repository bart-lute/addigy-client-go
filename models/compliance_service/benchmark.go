package compliance_service

type Benchmark struct {
	//compliance_rules	[compliance_service.ComplianceRule{...}]
	ComplianceRulesIds []string `json:"compliance_rules_ids"`
	CreatedDate        string   `json:"created_date"`
	Id                 string   `json:"id"`
	IsBaseline         bool     `json:"is_baseline"`
	MaximumOSVersion   string   `json:"maximum_os_version"`
	MinimumOSVersion   string   `json:"minimum_os_version"`
	Name               string   `json:"name"`
	OrganizationId     string   `json:"organization_id"`
	Source             string   `json:"source"`
	SourceBenchmarkId  string   `json:"source_benchmark_id"`
	TargetOS           string   `json:"target_os"`
	UpdatedDate        string   `json:"updated_date"`
}
