package compliance_service

import "github.com/bart-lute/addigy-client-go/models/entities"

type ComplianceRuleUpdateRequest struct {
	AgentRemediationScriptId string                 `json:"agent_remediation_script_id"`
	Description              string                 `json:"description"`
	FilterSets               []entities.AuditFilter `json:"filter_sets"`
	Id                       string                 `json:"id"`
	Name                     string                 `json:"name"`
	RemediationEnabled       bool                   `json:"remediation_enabled"`
}
