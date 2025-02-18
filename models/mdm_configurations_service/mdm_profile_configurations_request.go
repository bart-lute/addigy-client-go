package mdm_configurations_service

type MdmProfileConfigurationsRequest struct {
	ExcludedPayloadGroupIds []string `json:"excluded_payload_group_ids"`
	PayloadGroupIds         []string `json:"payload_group_ids"`
}
