package default_assets_service

type DefaultMdmConfigurationQuery struct {
	Ids          []string `json:"ids"`
	NameContains string   `json:"name_contains"`
}
