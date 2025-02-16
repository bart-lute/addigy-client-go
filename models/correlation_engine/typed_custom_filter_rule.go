package correlation_engine

import "github.com/bart-lute/addigy-client-go/models"

type TypedCustomFilterRule struct {
	Created        string                  `json:"created"`
	OrganizationId string                  `json:"organization_id"`
	Payload        models.CustomFilterRule `json:"payload"`
	RuleId         string                  `json:"rule_id"`
}
