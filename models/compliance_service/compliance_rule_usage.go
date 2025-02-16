package compliance_service

type ComplianceRuleUsage struct {
	Benchmarks []Benchmark `json:"benchmarks"`
}
