package devices

import "github.com/bart-lute/addigy-client-go/models/entities"

type DeviceFiltersResponse struct {
	Explanation  string                 `json:"explanation"`
	Filters      []entities.AuditFilter `json:"filters"`
	InvalidFacts int                    `json:"invalid_facts"`
	MissingFacts int                    `json:"missing_facts"`
}
