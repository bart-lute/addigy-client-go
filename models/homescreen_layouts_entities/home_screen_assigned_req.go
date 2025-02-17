package homescreen_layouts_entities

type HomeScreenAssignedReq struct {
	Assigned bool   `json:"assigned"`
	PolicyId string `json:"policyId"`
}
