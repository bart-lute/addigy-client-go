package instructions_service

import "encoding/json"

type CustomSoftwarePostRequest struct {
	Archived             bool                 `json:"archived"`
	BaseIdentifier       string               `json:"base_identifier"`
	Category             string               `json:"category"`
	Condition            string               `json:"condition"`
	Description          string               `json:"description"`
	Downloads            []Download           `json:"downloads"`
	Identifier           string               `json:"identifier"`
	InstallationScript   string               `json:"installation_script"`
	PredefinedConditions PredefinedConditions `json:"predefined_conditions"`
	Priority             json.Number          `json:"priority"`
	Profiles             []Profile            `json:"profiles"`
	RemoveScript         string               `json:"remove_script"`
	RunOnSuccess         bool                 `json:"run_on_success"`
	SoftwareIcon         Download             `json:"software_icon"`
	StatusOnSkipped      string               `json:"status_on_skipped"` // [ finished, failed ]
	UserEmail            string               `json:"user_email"`
	Version              string               `json:"version"`
}
