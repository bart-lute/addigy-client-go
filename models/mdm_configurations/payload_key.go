package mdm_configurations

type PayloadKey struct {
	Default              interface{}        `json:"default"`
	Description          string             `json:"description"`
	JsonIdentifier       string             `json:"json_identifier"`
	Keys                 []interface{}      `json:"keys"`
	Name                 string             `json:"name"`
	PlistIdentifier      string             `json:"plist_identifier"`
	RequiredFeatureFlags []string           `json:"required_feature_flags"`
	Type                 string             `json:"type"`
	UI                   UI                 `json:"ui"`
	Validators           []PayloadValidator `json:"validators"`
}
