package instructions_service

type ExamplePPPCProfile struct {
	AddigyPayloadType string                 `json:"addigy_payload_type"`
	Bundle            string                 `json:"bundle"`
	Custom            bool                   `json:"custom"`
	Event             []string               `json:"event"`
	FileId            string                 `json:"file_id"`
	Identifier        string                 `json:"identifier"`
	IdentifierType    string                 `json:"identifier_type"`
	Name              string                 `json:"name"`
	Path              string                 `json:"path"`
	Permissions       ExamplePPPCPermissions `json:"permissions"`
	Requirements      string                 `json:"requirements"`
	Signature         string                 `json:"signature"`
}
