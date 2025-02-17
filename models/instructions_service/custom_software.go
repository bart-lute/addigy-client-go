package instructions_service

import "encoding/json"

type CustomSoftware struct {
	Archived                    bool                 `json:"archived"`
	BaseIdentifier              string               `json:"base_identifier"`
	Category                    string               `json:"category"`
	Condition                   string               `json:"condition"`
	Description                 string               `json:"description"`
	Downloads                   []Download           `json:"downloads"`
	FactIdentifier              string               `json:"fact_identifier"`
	Identifier                  string               `json:"identifier"`
	Install                     bool                 `json:"install"`
	InstallationScript          string               `json:"installation_script"`
	InstructionId               string               `json:"instruction_id"`
	IsOnboardingConfig          bool                 `json:"is_onboarding_config"`
	Label                       string               `json:"label"`
	Name                        string               `json:"name"`
	OrganizationId              string               `json:"organization_id"`
	PolicyRestricted            bool                 `json:"policy_restricted"`
	PredefinedConditions        PredefinedConditions `json:"predefined_conditions"`
	PricePerDevice              json.Number          `json:"price_per_device"`
	Priority                    json.Number          `json:"priority"`
	Profiles                    []Profile            `json:"profiles"`
	Provider                    string               `json:"provider"`
	Public                      bool                 `json:"public"`
	PublicSoftwareInstructionId string               `json:"public_software_instruction_id"`
	RemoveScript                string               `json:"remove_script"`
	RunOnSuccess                bool                 `json:"run_on_success"`
	SetGroupName                bool                 `json:"set_group_name"`
	SoftwareIcon                Download             `json:"software_icon"`
	StatusOnSkipped             string               `json:"status_on_skipped"`
	TccVersion                  int                  `json:"tcc_version"`
	Type                        string               `json:"type"`
	UserEmail                   string               `json:"user_email"`
	Version                     string               `json:"version"`
}
