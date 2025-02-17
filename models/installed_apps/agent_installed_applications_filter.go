package installed_apps

type AgentInstalledApplicationsFilter struct {
	AgentIds []string `json:"agent_ids"`
	Names    []string `json:"names"`
}
