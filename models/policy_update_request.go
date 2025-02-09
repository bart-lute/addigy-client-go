package models

type PolicyUpdateRequest struct {
	Color    string `json:"color"`
	Icon     string `json:"icon"`
	Name     string `json:"name"`
	PolicyId string `json:"policy_id"`
}
