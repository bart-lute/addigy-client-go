package device_users_service

type EndUserDeviceQuery struct {
	AgentIds   []string `json:"agent_ids"`
	EndUserIds []string `json:"end_user_ids"`
}
