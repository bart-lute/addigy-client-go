package models

type PolicyParentUpdateRequest struct {
	ParentPolicyId string `json:"parent_policy_id"`
	PolicyId       string `json:"policy_id"`
}
