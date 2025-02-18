package mdm_service_devices

type MdmCommandRequest struct {
	AgentId    string `json:"agent_id"`
	DeviceUuid string `json:"device_uuid"`
}
