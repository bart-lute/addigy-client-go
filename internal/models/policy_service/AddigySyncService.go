package policy_service

type AddigySyncService struct {
	AdminGroups     []string `json:"admin_groups"`
	AllowLocalLogin bool     `json:"allow_local_login"`
	APIKey          string   `json:"api_key"`
	BackgroundImage Image    `json:"background_image"`
	Domain          string   `json:"domain"`
	IsAdmin         bool     `json:"is_admin"`
	LoginLogo       Image    `json:"login_logo"`
}
