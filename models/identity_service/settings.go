package identity_service

type Settings struct {
	ActiveService      string   `json:"active_service"`
	AwaitConfiguration bool     `json:"await_configuration"`
	ConfigVersion      int      `json:"config_version"`
	Services           Services `json:"services"`
}
