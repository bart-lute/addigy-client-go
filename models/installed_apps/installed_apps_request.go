package installed_apps

type InstalledAppsRequest struct {
	AgentIds      []string `json:"agent_ids"`
	Limit         int      `json:"limit"`
	Skip          int      `json:"skip"`
	SortDirection string   `json:"sort_direction"` // asc,desc
	SortField     string   `json:"sort_field"`
}
