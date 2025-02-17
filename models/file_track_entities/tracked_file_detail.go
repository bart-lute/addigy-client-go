package file_track_entities

type TrackedFileDetail struct {
	FeatureName string `json:"feature_name"`
	FeatureType string `json:"feature_type"`
	FileId      string `json:"file_id"`
	ItemId      string `json:"item_id"`
	ItemName    string `json:"item_name"`
}
