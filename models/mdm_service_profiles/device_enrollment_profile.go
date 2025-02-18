package mdm_service_profiles

import "encoding/json"

type DeviceEnrollmentProfile struct {
	HasRemovalPasscode       bool               `json:"has_removal_passcode"`
	InstalledPayloads        []InstalledPayload `json:"installed_payloads"`
	IsEncrypted              bool               `json:"is_encrypted"`
	LastUpdated              string             `json:"last_updated"`
	PayloadDescription       string             `json:"payload_description"`
	PayloadDisplayName       string             `json:"payload_display_name"`
	PayloadIdentifier        string             `json:"payload_identifier"`
	PayloadOrganization      string             `json:"payload_organization"`
	PayloadRemovalDisallowed bool               `json:"payload_removal_disallowed"`
	PayloadUuid              string             `json:"payload_uuid"`
	PayloadVersion           interface{}        `json:"payload_version"`
	Priority                 json.Number        `json:"priority"`
	Udid                     string             `json:"udid"`
	UserId                   string             `json:"user_id"`
}
