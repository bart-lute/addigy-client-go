package alert_entities

import "encoding/json"

type AlertPost struct {
	Category           string            `json:"category"`
	Emails             []string          `json:"emails"`
	Fact               string            `json:"fact"`
	FactIdentifier     string            `json:"fact_identifier"`
	HasScript          bool              `json:"has_script"`
	Instructions       map[string]string `json:"instructions"` // Revisit later
	IsInBlueprint      bool              `json:"is_in_blueprint"`
	Level              string            `json:"level"`
	MaxValue           json.Number       `json:"max_value"`
	MinValue           json.Number       `json:"min_alue"`
	Name               string            `json:"name"`
	RemediationEnabled bool              `json:"remediation_enabled"`
	RemediationTime    int               `json:"remediation_time"`
	ScriptId           string            `json:"script_id"`
	Selector           string            `json:"selector"`
	SendTicket         bool              `json:"send_ticket"`
	Type               string            `json:"type"`
	Value              map[string]string `json:"value"` // Revisit later
	ValueType          string            `json:"value_type"`
	Version            int               `json:"version"`
}
