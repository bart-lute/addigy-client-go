package compliance_service

type DeviceComplianceBenchmarkStatus struct {
	AgentId        string       `json:"agent_id"`
	BenchmarkId    string       `json:"benchmark_id"`
	FailedRules    []FailedRule `json:"failed_rules"`
	IsCompliant    bool         `json:"is_compliant"`
	LastUpdated    string       `json:"last_updated"`
	OrganizationId string       `json:"organization_id"`
}
