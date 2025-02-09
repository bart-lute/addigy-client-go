package alert_entities

import "encoding/json"

type ScheduledAlert struct {
	Id               string            `json:"_id"`
	Category         string            `json:"category"`
	Description      string            `json:"description"`
	Emails           []string          `json:"emails"`
	Fact             string            `json:"fact"`
	FactIdentifier   string            `json:"fact_identifier"`
	HasScript        bool              `json:"has_script"`
	Instructions     map[string]string `json:"instructions"` // Revisit later
	Level            string            `json:"level"`
	MaxValue         json.Number       `json:"max_value"`
	MinValue         json.Number       `json:"min_value"`
	Name             string            `json:"name"`
	Orgid            string            `json:"orgid"`
	PolicyRestricted bool              `json:"policy_restricted"`
	Provider         string            `json:"provider"`
	Remenabled       bool              `json:"remenabled"`
	Remtime          int               `json:"remtime"`
	Script           []string          `json:"script"`
	ScriptId         string            `json:"script_id"`
	Selector         string            `json:"selector"`
	SendTicket       bool              `json:"send_ticket"`
	Type             string            `json:"type"`
	Value            map[string]string `json:"value"`
	Valuetype        string            `json:"valuetype"`
}
