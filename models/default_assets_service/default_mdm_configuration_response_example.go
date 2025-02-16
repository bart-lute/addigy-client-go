package default_assets_service

type DefaultMdmConfigurationResponseExample struct {
	AddigyPayloadType    string `json:"addigy_payload_type"`
	AddigyPayloadVersion int    `json:"addigy_payload_version"`
	PayloadDisplayName   string `json:"payload_display_name"`
	PayloadEnabled       bool   `json:"payload_enabled"`
	PayloadGroupId       string `json:"payload_group_id"`
	PayloadIdentifier    string `json:"payload_identifier"`
	PayloadPriority      int    `json:"payload_priority"`
	PayloadType          string `json:"payload_type"`
	PayloadUuid          string `json:"payload_uuid"`
	PayloadVersion       int    `json:"payload_version"`
	PolicyRestricted     bool   `json:"policy_restricted"`
}
