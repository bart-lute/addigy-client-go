package scripts_service

type DeviceScriptAssignmentResponse struct {
	AgentId  string `json:"agent_id"`
	Device   Device `json:"device"`
	Script   Script `json:"script"`
	ScriptId string `json:"script_id"`
}
