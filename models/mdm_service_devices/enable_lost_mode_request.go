package mdm_service_devices

type EnableLostModeRequest struct {
	AgentId     string `json:"agent_id"`
	DeviceUuid  string `json:"device_uuid"`
	Footnote    string `json:"footnote"`
	Message     string `json:"message"`
	PhoneNumber string `json:"phone_number"`
}
