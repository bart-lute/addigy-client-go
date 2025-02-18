package maintenance_entities

type Maintenance struct {
	_Id                     string               `json:"_id"`
	Day                     string               `json:"day"`
	Description             string               `json:"description"`
	Enabled                 bool                 `json:"enabled"`
	ExpectedRemediationTime int                  `json:"expected_remediation_time"`
	Frequency               string               `json:"frequency"`
	IsInBlueprint           bool                 `json:"is_in_blueprint"`
	LocalTime               bool                 `json:"local_time"`
	MaxTryCount             int                  `json:"max_try_count"`
	Name                    string               `json:"name"`
	Policies                []string             `json:"policies"`
	PolicyRestricted        bool                 `json:"policy_restricted"`
	PromptUser              bool                 `json:"prompt_user"`
	ScheduledTime           string               `json:"scheduled_time"`
	Scripts                 []InstructionCommand `json:"scripts"`
	Source                  string               `json:"source"`
	SourceId                string               `json:"source_id"`
	TimeoutSeconds          int                  `json:"timeout_seconds"`
	Version                 int                  `json:"version"`
}
