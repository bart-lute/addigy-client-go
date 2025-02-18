package mdm_configurations

type MDMProfiles struct {
	AddigyPayloadType    string `json:"addigy_payload_type"`
	AddigyPayloadVersion int    `json:"addigy_payload_version"`
	PayloadDisplayName   string `json:"payload_display_name"`
	PayloadGroupId       string `json:"payload_group_id"`
	PayloadIdentifier    string `json:"payload_identifier"`
	PayloadType          string `json:"payload_type"`
	PayloadUuid          string `json:"payload_uuid"`
	PolicyRestricted     bool   `json:"policy_restricted"`
}
