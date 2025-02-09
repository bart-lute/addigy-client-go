package policy_service

import "encoding/json"

type SystemUpdatesSettings struct {
	ForceUnattended      bool        `json:"force_unattended"`
	ForceUpdate          bool        `json:"force_update"`
	Restart              bool        `json:"restart"`
	RestartAfterAttempts int         `json:"restart_after_attempts"`
	RestartHard          bool        `json:"restart_hard"`
	RestartPromptPeriod  json.Number `json:"restart_prompt_period"`
	RestartPromptUser    bool        `json:"restart_prompt_user"`
}
