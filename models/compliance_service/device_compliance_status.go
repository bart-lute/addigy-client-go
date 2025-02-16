package compliance_service

type DeviceComplianceStatus struct {
	AgentId     string `json:"agent_id"`
	IsCompliant bool   `json:"is_compliant"`
}
