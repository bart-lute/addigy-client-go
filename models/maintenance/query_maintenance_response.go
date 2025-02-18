package maintenance

import (
	"github.com/bart-lute/addigy-client-go/models/maintenance_entities"
	"github.com/bart-lute/addigy-client-go/models/response_entities"
)

type QueryMaintenanceResponse struct {
	Items    []maintenance_entities.Maintenance `json:"items"`
	Metadata response_entities.Metadata         `json:"metadata"`
}
