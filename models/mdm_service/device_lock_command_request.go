package mdm_service

type DeviceLockCommandRequest struct {
	AgentId     string `json:"agent_id"`
	Message     string `json:"message"`
	PhoneNumber string `json:"phone_number"`
	Pin         string `json:"pin"`
	Uuid        string `json:"uuid"`
}
