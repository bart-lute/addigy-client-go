package mdm_configurations

type PayloadManifest struct {
	AddigyPayloadType   string                `json:"addigy_payload_type"`
	Availability        PayloadAvailability   `json:"availability"`
	HasManifest         bool                  `json:"has_manifest"`
	IconPath            IconPath              `json:"icon_path"`
	Keys                []PayloadKey          `json:"keys"`
	PayloadDescription  string                `json:"payload_description"`
	PayloadName         string                `json:"payload_name"`
	PayloadObject       string                `json:"payload_object"`
	PayloadType         string                `json:"payload_type"`
	SupportedOsVersions []SupportedOsVersions `json:"supported_os_versions"`
}
