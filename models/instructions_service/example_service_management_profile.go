package instructions_service

type Rule struct {
	RuleType  string `json:"rule_type"`
	RuleValue string `json:"rule_value"`
}

type ExampleServiceManagementProfile struct {
	AddigyPayloadType string   `json:"addigy_payload_type"`
	Bundle            string   `json:"bundle"`
	Custom            bool     `json:"custom"`
	Events            []string `json:"events"`
	FileId            string   `json:"file_id"`
	Identifier        string   `json:"identifier"`
	IdentifierType    string   `json:"identifier_type"`
	Name              string   `json:"name"`
	Path              string   `json:"path"`
	Requirements      string   `json:"requirements"`
	Rules             []Rule   `json:"rules"`
	Signature         string   `json:"signature"`
}
