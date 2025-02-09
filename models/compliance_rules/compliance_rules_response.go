package compliance_rules

import (
	"github.com/bart-lute/addigy-client-go/models/compliance_service"
	"github.com/bart-lute/addigy-client-go/models/response_entities"
)

type ComplianceRulesResponse struct {
	Items    []compliance_service.ComplianceRule `json:"items"`
	Metadata response_entities.Metadata          `json:"metadata"`
}
