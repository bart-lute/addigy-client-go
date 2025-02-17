package instructions_service

type ExampleSystemExtensionProfile struct {
	AddigyPayloadType         string            `json:"addigy_payload_type"`
	AllowedSystemExtensions   map[string]string `json:"allowed_system_extensions"` // example: OrderedMap { "BJ4HAAB9B3": "us.zoom.xos" }
	Bundle                    string            `json:"bundle"`
	Custom                    bool              `json:"custom"`
	FileId                    string            `json:"file_id"`
	Identifier                string            `json:"identifier"`
	IdentifierType            string            `json:"identifier_type"`
	Name                      string            `json:"name"`
	Path                      string            `json:"path"`
	RemovableSystemExtensions map[string]string `json:"removable_system_extensions"` // example: OrderedMap { "BJ4HAAB9B3": "us.zoom.xos" }
	Requirements              string            `json:"requirements"`
	Signature                 string            `json:"signature"`
}
