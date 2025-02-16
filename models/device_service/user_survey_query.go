package device_service

type UserSurveyQuery struct {
	AgentIds       []string `json:"agent_ids"`
	EndTimestamp   int      `json:"end_timestamp"`
	StartTimestamp int      `json:"start_timestamp"`
}
