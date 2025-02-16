package device_service

type UserSurvey struct {
	Agentid        string `json:"agentid"`
	Happy          bool   `json:"happy"`
	OrganizationId string `json:"organization_id"`
}
