package instructions_service

type Profile struct {
	AddigyPayloadType            string                  `json:"addigy_payload_type"` // [ com.addigy.policy.smart-software.servicemanagement.com.apple.servicemanagement, com.addigy.policy.smart-software.system-extension-policy.com.apple.system-extension-policy, com.addigy.policy.smart-software.pppc.com.addigy.pppc ]
	AllowUserOverrides           bool                    `json:"allow_user_overrides"`
	AllowedSystemExtensions      map[string]string       `json:"allowed_system_extensions"`
	AllowedSystemExtensionsTypes map[string]string       `json:"allowed_system_extensions_types"`
	AllowedTeamIdentifiers       []string                `json:"allowed_team_identifiers"`
	Bundle                       string                  `json:"bundle"`
	Custom                       bool                    `json:"custom"`
	Events                       []ProfileEvent          `json:"events"`
	FileId                       string                  `json:"file_id"`
	Identifier                   string                  `json:"identifier"`
	IdentifierType               string                  `json:"identifier_type"`
	Name                         string                  `json:"name"`
	Path                         string                  `json:"path"`
	Permissions                  ProfilePermissions      `json:"permissions"`
	RemovableSystemExtensions    map[string]string       `json:"removable_system_extensions"`
	Requirements                 string                  `json:"requirements"`
	Rules                        []ServiceManagementRule `json:"rules"`
	Signature                    string                  `json:"signature"`
}
