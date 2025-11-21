package models

import "github.com/bart-lute/addigy-client-go/internal/models/policy_service"

type Policy struct {
	AddigySync                policy_service.AddigySync            `json:"addigy_sync"`
	AgentPath                 string                               `json:"agent_path"`
	AgentVersion              string                               `json:"agent_version"`
	AutotaskAccountID         int                                  `json:"autotask_account_id"`
	CollectorSettings         policy_service.CollectorSettings     `json:"collector_settings"`
	Color                     string                               `json:"color"`
	ConnectwiseAccountID      int                                  `json:"connectwise_account_id"`
	CreationTime              int                                  `json:"creation_time"`
	DownloadPath              string                               `json:"download_path"`
	Icon                      string                               `json:"icon"`
	IgnoreUpdates             bool                                 `json:"ignore_updates"`
	Instructions              []string                             `json:"instructions"`
	ItglueAccountID           string                               `json:"itglue_account_id"`
	LastDeployed              string                               `json:"last_deployed"`
	Name                      string                               `json:"name"`
	Orgid                     string                               `json:"orgid"`
	Parent                    string                               `json:"parent"`
	PolicyID                  string                               `json:"policyId"`
	PrebuiltAppSettings       policy_service.PrebuiltAppSettings   `json:"prebuilt_app_settings"`
	Schedules                 map[string]policy_service.Schedule   `json:"schedules"`
	SelfServiceInstructionIds map[string]bool                      `json:"self_service_instruction_ids"`
	SplashtopSettings         policy_service.SplashtopSettings     `json:"splashtop_settings"`
	SSHSettings               policy_service.SshSettings           `json:"ssh_settings"`
	SystemUpdatesSettings     policy_service.SystemUpdatesSettings `json:"system_updates_settings"`
	VncSettings               policy_service.VncSettings           `json:"vnc_settings"`
}
