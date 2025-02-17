package identity_service

type SettingsStore struct {
	ActiveService      string   `json:"active_service"`
	AwaitConfiguration bool     `json:"await_configuration"`
	Services           Services `json:"services"`
}
