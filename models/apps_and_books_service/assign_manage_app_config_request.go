package apps_and_books_service

type AssignManageAppConfigRequest struct {
	BundleId   string `json:"bundle_id"`
	IsAssigned bool   `json:"is_assigned"`
	LocationId string `json:"location_id"`
}
