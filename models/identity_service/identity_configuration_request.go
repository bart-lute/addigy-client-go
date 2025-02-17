package identity_service

type IdentityConfigurationRequest struct {
	Ids       []string `json:"ids"`
	PolicyIds []string `json:"policy_ids"`
}
