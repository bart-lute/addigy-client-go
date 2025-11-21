package models

import (
    policy_service2 "github.com/bart-lute/addigy-client-go/models/policy_service"
)

type Policy struct {
    AddigySync                policy_service2.AddigySync            `json:"addigy_sync"`
    AgentPath                 string                                `json:"agent_path"`
    AgentVersion              string                                `json:"agent_version"`
    AutotaskAccountID         int                                   `json:"autotask_account_id"`
    CollectorSettings         policy_service2.CollectorSettings     `json:"collector_settings"`
    Color                     string                                `json:"color"`
    ConnectwiseAccountID      int                                   `json:"connectwise_account_id"`
    CreationTime              int                                   `json:"creation_time"`
    DownloadPath              string                                `json:"download_path"`
    Icon                      string                                `json:"icon"`
    IgnoreUpdates             bool                                  `json:"ignore_updates"`
    Instructions              []string                              `json:"instructions"`
    ItglueAccountID           string                                `json:"itglue_account_id"`
    LastDeployed              string                                `json:"last_deployed"`
    Name                      string                                `json:"name"`
    Orgid                     string                                `json:"orgid"`
    Parent                    string                                `json:"parent"`
    PolicyID                  string                                `json:"policyId"`
    PrebuiltAppSettings       policy_service2.PrebuiltAppSettings   `json:"prebuilt_app_settings"`
    Schedules                 map[string]policy_service2.Schedule   `json:"schedules"`
    SelfServiceInstructionIds map[string]bool                       `json:"self_service_instruction_ids"`
    SplashtopSettings         policy_service2.SplashtopSettings     `json:"splashtop_settings"`
    SSHSettings               policy_service2.SshSettings           `json:"ssh_settings"`
    SystemUpdatesSettings     policy_service2.SystemUpdatesSettings `json:"system_updates_settings"`
    VncSettings               policy_service2.VncSettings           `json:"vnc_settings"`
}
