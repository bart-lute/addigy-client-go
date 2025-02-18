package mdm_service_devices

type LostModeRequest struct {
	AgentId    string `json:"agent_id"`
	DeviceUuid string `json:"device_uuid"`
}
