package maintenance_entities

import "encoding/json"

type StagedMaintenance struct {
	Id               string      `json:"_id"`
	Day              string      `json:"day"`
	Description      string      `json:"description"`
	Enabled          bool        `json:"enabled"`
	Frequency        string      `json:"frequency"`    // example: Monthly
	Instructions     []string    `json:"instructions"` // Not sure
	IsInBlueprint    bool        `json:"is_in_blueprint"`
	Jobname          string      `json:"jobname"`
	Jobtime          json.Number `json:"jobtime"`
	LocalTime        bool        `json:"local_time"`
	Maxtrycount      int         `json:"maxtrycount"`
	Policies         []string    `json:"policies"`
	PolicyRestricted bool        `json:"policy_restricted"`
	Promptuser       bool        `json:"promptuser"`
	Scheduledtime    string      `json:"scheduledtime"`
	Source           string      `json:"source"`
	SourceId         string      `json:"source_id"`
	TimeoutSeconds   int         `json:"timeout_seconds"`
	Version          int         `json:"version"`
}
