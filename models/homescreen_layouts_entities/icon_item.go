package homescreen_layouts_entities

type IconItem struct {
	BundleId    string        `json:"bundle_id"`
	DisplayName string        `json:"display_name"`
	Pages       []interface{} `json:"pages"` // This needs to get fixed later on
	Type        string        `json:"type"`
	Url         string        `json:"url"`
}
