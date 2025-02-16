package default_assets_entities

type DefaultMaintenance struct {
	Day              string               `json:"day"`
	Enabled          bool                 `json:"enabled"`
	Frequency        string               `json:"frequency"`
	Id               string               `json:"id"`
	Instructions     []InstructionCommand `json:"instructions"`
	JobTime          int                  `json:"job_time"`
	MaxTryCount      int                  `json:"max_try_count"`
	Name             string               `json:"name"`
	PromptUser       bool                 `json:"prompt_user"`
	ScheduledTime    string               `json:"scheduled_time"`
	ShortDescription string               `json:"short_description"`
}
