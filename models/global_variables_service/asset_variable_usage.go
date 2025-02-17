package global_variables_service

type AssetVariableUsage struct {
	AssetIds       []string `json:"asset_ids"`
	AssetType      string   `json:"asset_type"`
	OrganizationId string   `json:"organization_id"`
	VariableKey    string   `json:"variable_key"`
}
