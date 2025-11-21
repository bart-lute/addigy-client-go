package policy_service

type CollectorSettings struct {
	CollectSystemAppsUsage bool `json:"collect_system_apps_usage"`
	CollectWebAppsUsage    bool `json:"collect_web_apps_usage"`
}
