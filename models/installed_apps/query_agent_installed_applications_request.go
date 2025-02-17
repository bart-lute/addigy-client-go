package installed_apps

type QueryAgentInstalledApplicationsRequest struct {
	Page          int                              `json:"page"`
	PerPage       int                              `json:"per_page"`
	Query         AgentInstalledApplicationsFilter `json:"query"`
	SortDirection string                           `json:"sort_direction"`
	SortField     string                           `json:"sort_field"`
}
