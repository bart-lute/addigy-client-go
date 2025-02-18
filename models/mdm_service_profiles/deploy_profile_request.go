package mdm_service_profiles

type DeployProfileRequest struct {
	DevicesUserIds  map[string]string `json:"devices_user_ids"` // ???
	DevicesUuid     []string          `json:"devices_uuid"`
	PayloadsGroupId string            `json:"payloads_group_id"`
}
