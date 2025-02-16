package default_assets_service

type DefaultSelfServiceConfigurationQuery struct {
	Ids          []string `json:"ids"`
	NameContains string   `json:"name_contains"`
}
