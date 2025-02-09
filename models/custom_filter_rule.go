package models

import "github.com/bart-lute/addigy-client-go/models/entities"

type CustomFilterRule struct {
	AutoRemove bool                   `json:"auto_remove"`
	Disabled   bool                   `json:"disabled"`
	Filters    []entities.AuditFilter `json:"filters"`
	PolicyId   string                 `json:"policy_id"`
}
