package policy_service

type VncSettings struct {
	Enabled                bool `json:"enabled"`
	RequireUserPermissions bool `json:"require_user_permissions"`
}
