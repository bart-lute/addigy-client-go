package policy_service

// Note: Based on actual returned data, not on API model

type PolicyMdmConfiguration struct {
	Orgid           string `json:"orgid"`
	ConfigurationID string `json:"configuration_id"`
	PolicyID        string `json:"policy_id"`
}
