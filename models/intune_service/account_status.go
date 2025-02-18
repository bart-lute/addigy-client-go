package intune_service

type AccountStatus struct {
	Enabled        bool   `json:"enabled"`
	OrganizationId string `json:"organization_id"`
	PolicyId       string `json:"policy_id"`
}
