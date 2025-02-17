package homescreen_layouts_entities

type HomeScreenCreateReq struct {
	Assigned   bool         `json:"assigned"`
	DeviceType string       `json:"device_type"`
	Dock       []IconItem   `json:"dock"`
	Pages      [][]IconItem `json:"pages"` // Check later on. Does this even make sense?
	PolicyId   string       `json:"policy_id"`
}
