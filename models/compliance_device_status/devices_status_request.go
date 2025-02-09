package compliance_device_status

type DevicesStatusRequest struct {
	AgentIds []string `json:"agent_ids"`
}
