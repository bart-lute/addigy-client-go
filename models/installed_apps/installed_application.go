package installed_apps

type InstalledApplication struct {
	AgentId            string `json:"agent_id"`
	HasUpdateAvailable bool   `json:"has_update_available"`
	Identifier         string `json:"identifier"`
	LastUpdated        string `json:"last_updated"`
	Name               string `json:"name"`
	Orgid              string `json:"orgid"`
	ShortVersion       string `json:"short_version"`
	Udid               string `json:"udid"`
	Version            string `json:"version"`
}
