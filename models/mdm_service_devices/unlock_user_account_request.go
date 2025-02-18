package mdm_service_devices

type UnlockUserAccountRequest struct {
	AgentId    string `json:"agent_id"`
	DeviceUuid string `json:"device_uuid"`
	Username   string `json:"username"`
}
