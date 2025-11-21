package policy_service

type AddigySync struct {
	ActiveService      string                       `json:"active_service"`
	AwaitConfiguration bool                         `json:"await_configuration"`
	ConfigVersion      int                          `json:"config_version"`
	Services           map[string]AddigySyncService `json:"services"`
}
