package default_assets_service

type DefaultMaintenanceQuery struct {
	Ids          []string `json:"ids"`
	NameContains string   `json:"name_contains"`
}
