package device_users_service

type EndUserDevice struct {
	AgentId        string `json:"agent_id"`
	EndUserId      string `json:"end_user_id"`
	OrganizationId string `json:"organization_id"`
}
