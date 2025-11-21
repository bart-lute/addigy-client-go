package policy_service

// NOTE: The API Documentation Model does not match the output
// Below is the Struct that is based on the actual output, not on the model used in the API

type MdmPayloadJSONResult struct {
	AddigyPayloadType    string  `json:"addigy_payload_type"`
	AddigyPayloadVersion int     `json:"addigy_payload_version"`
	HasManifest          bool    `json:"has_manifest"`
	IsInBlueprint        bool    `json:"is_in_blueprint"`
	Orgid                string  `json:"orgid"`
	PayloadDisplayName   string  `json:"payload_display_name"`
	PayloadEnabled       bool    `json:"payload_enabled"`
	PayloadGroupID       string  `json:"payload_group_id"`
	PayloadIdentifier    string  `json:"payload_identifier"`
	PayloadPriority      float32 `json:"payload_priority"`
	PayloadType          string  `json:"payload_type"`
	PayloadUUID          string  `json:"payload_uuid"`
	PayloadVersion       int     `json:"payload_version"`
	PolicyRestricted     bool    `json:"policy_restricted"`
	SourceID             string  `json:"source_id"`
	Version              int     `json:"version"`
}
