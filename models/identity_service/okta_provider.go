package identity_service

type OktaProvider struct {
	AllowLocalLogin       bool   `json:"allow_local_login"`
	AllowRevertLogin      bool   `json:"allow_revert_login"`
	AllowSyncUsers        bool   `json:"allow_sync_users"`
	ApiToken              string `json:"api_token"`
	BackgroundImage       Image  `json:"background_image"`
	ClientId              string `json:"client_id"`
	CollectUserAttributes bool   `json:"collect_user_attributes"`
	Domain                string `json:"domain"`
	HasApiManagement      bool   `json:"has_api_management"`
	IsAdmin               bool   `json:"is_admin"`
	LoginLogo             Image  `json:"login_logo"`
	RedirectUri           string `json:"redirect_uri"`
}
