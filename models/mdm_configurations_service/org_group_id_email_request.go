package mdm_configurations_service

type OrgGroupIdEmailRequest struct {
	GroupId   string   `json:"group_id"`
	PolicyIds []string `json:"policy_ids"`
}
