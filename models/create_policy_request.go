package models

type CreatePolicyRequest struct {
	Color          string `json:"color"`
	Icon           string `json:"icon"`
	Name           string `json:"name"`
	ParentPolicyId string `json:"parent_policy_id"`
}
