package mdm_configurations

type MDMProfiles struct {
	AddigyPayloadType    string `json:"addigy_payload_type"`
	AddigyPayloadVersion int    `json:"addigy_payload_version"`
	PayloadDisplayName   string `json:"payload_display_name"`
	PayloadGroupID       string `json:"payload_group_id"`
	PayloadIdentifier    string `json:"payload_identifier"`
	PayloadType          string `json:"payload_type"`
	PayloadUUID          string `json:"payload_uuid"`
	PayloadRestricted    int    `json:"payload_restricted"`
}
