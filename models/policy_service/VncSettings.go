package policy_service

type VncSettings struct {
	Enabled               bool `json:"enabled"`
	RequireUserPermission bool `json:"require_user_permission"`
}
