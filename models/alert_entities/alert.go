package alert_entities

import "encoding/json"

type Alert struct {
	Category           string            `json:"category"`
	Description        string            `json:"description"`
	Emails             []string          `json:"emails"`
	FactIdentifier     string            `json:"fact_identifier"`
	HasScript          bool              `json:"has_script"`
	Id                 string            `json:"id"`
	Instructions       map[string]string `json:"instructions"` //Type is a bit unclear at the time of writing this
	IsInBlueprint      bool              `json:"is_in_blueprint"`
	Level              string            `json:"level"`
	MaxValue           json.Number       `json:"max_value"`
	MinValue           json.Number       `json:"min_value"`
	Name               string            `json:"name"`
	PolicyRestricted   bool              `json:"policy_restricted"`
	Provider           string            `json:"provider"`
	RemediationEnabled bool              `json:"remediation_enabled"`
	RemediationTime    int               `json:"remediation_time"`
	Script             []string          `json:"script"`
	ScriptId           string            `json:"script_id"`
	Selector           string            `json:"selector"`
	SendTicket         bool              `json:"send_ticket"`
	Source             string            `json:"source"`
	SourceId           string            `json:"source_id"`
	Type               string            `json:"type"`
	Value              map[string]string `json:"value"` //Type is a bit unclear at the time of writing this
	ValueType          string            `json:"value_type"`
	Version            int               `json:"version"`
}
