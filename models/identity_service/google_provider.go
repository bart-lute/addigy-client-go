package identity_service

type GoogleProvider struct {
	AllowLocalLogin       bool   `json:"allow_local_login"`
	AllowRevertLogin      bool   `json:"allow_revert_login"`
	AllowSyncUsers        bool   `json:"allow_sync_users"`
	BackgroundImage       Image  `json:"background_image"`
	ClientId              string `json:"client_id"`
	ClientSecret          string `json:"client_secret"`
	CollectUserAttributes bool   `json:"collect_user_attributes"`
	IsAdmin               bool   `json:"is_admin"`
	LoginLogo             Image  `json:"login_logo"`
	RedirectUri           string `json:"redirect_uri"`
}
