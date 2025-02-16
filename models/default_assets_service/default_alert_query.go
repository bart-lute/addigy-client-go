package default_assets_service

type DefaultAlertQuery struct {
	Ids          []string `json:"ids"`
	NameContains string   `json:"name_contains"`
}
