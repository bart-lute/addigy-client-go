package apps_and_books_service

type ManagedAppConfigurationRequest struct {
	AssetId       string                   `json:"asset_id"`
	BundleId      string                   `json:"bundle_id"`
	Configuration map[string]Configuration `json:"configuration"`
	IsAssigned    bool                     `json:"is_assigned"`
	LocationId    string                   `json:"location_id"`
}
