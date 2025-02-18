package maintenance_entities

type MaintenancePut struct {
	Day                     string               `json:"day"`
	Description             string               `json:"description"`
	Enabled                 bool                 `json:"enabled"`
	ExpectedRemediationTime int                  `json:"expected_remediation_time"`
	Frequency               string               `json:"frequency"`
	Id                      string               `json:"id"`
	IsInBlueprint           bool                 `json:"is_in_blueprint"`
	LocalTime               bool                 `json:"local_time"`
	MaxTryCount             int                  `json:"max_try_count"`
	Name                    string               `json:"name"`
	PromptUser              bool                 `json:"prompt_user"`
	ScheduledTime           string               `json:"scheduled_time"`
	Scripts                 []InstructionCommand `json:"scripts"`
	TimeoutSeconds          int                  `json:"timeout_seconds"`
	Version                 int                  `json:"version"`
}
