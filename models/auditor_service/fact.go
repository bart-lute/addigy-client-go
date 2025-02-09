package auditor_service

type Fact struct {
	CommunityFactId  string          `json:"community_fact_id"`
	CommunityVersion int             `json:"community_version"`
	Identifier       string          `json:"identifier"`
	InstructionId    string          `json:"instruction_id"`
	Name             string          `json:"name"`
	Notes            string          `json:"notes"`
	Orgid            string          `json:"orgid"`
	OsArchitectures  OSArchitectures `json:"os_architectures"`
	Provider         string          `json:"provider"`
	ReturnType       string          `json:"return_type"`
	Source           string          `json:"source"`
	Version          int             `json:"version"`
}
