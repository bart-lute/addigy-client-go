package device_enties

import "github.com/bart-lute/addigy-client-go/models/response_entities"

type DeviceAuditResponse struct {
	Items    []DeviceAudit              `json:"items"`
	Metadata response_entities.Metadata `json:"metadata"`
}
