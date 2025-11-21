package policy_service

type SshSettings struct {
	Enabled               bool `json:"enabled"`
	RequireUserPermission bool `json:"require_user_permission"`
}
