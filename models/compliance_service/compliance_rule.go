package compliance_service

import "github.com/bart-lute/addigy-client-go/models/entities"

type ComplianceRule struct {
	AgentRemediationScriptId string                 `json:"agent_remediation_script_id"`
	CreatedDate              string                 `json:"created_date"`
	Description              string                 `json:"description"`
	FilterSets               []entities.AuditFilter `json:"filter_sets"`
	Id                       string                 `json:"id"`
	Name                     string                 `json:"name"`
	RemediationEnabled       bool                   `json:"remediation_enabled"`
	Severity                 string                 `json:"severity"`
	UpdatedDate              string                 `json:"updated_date"`
}
