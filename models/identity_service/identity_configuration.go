package identity_service

type IdentityConfiguration struct {
	AddigySync Settings `json:"addigy_sync"`
	Id         string   `json:"id"`
	OrgId      string   `json:"org_id"`
	PolicyId   string   `json:"policy_id"`
}
