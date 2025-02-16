package compliance_service

type PreBuiltBenchmarkCloneRequest struct {
	BenchmarkId        string   `json:"benchmark_id"`
	ComplianceRulesIds []string `json:"compliance_rules_ids"` // Minimum Items: 1
	Name               string   `json:"name"`
}
