package alert_entities

import "github.com/bart-lute/addigy-client-go/models/response_entities"

type AlertsScheduleResponse struct {
	AlertPolicies map[string]string          `json:"alert_policies"`
	Alerts        response_entities.Response `json:"alerts"`
	//StagedAlerts
}
