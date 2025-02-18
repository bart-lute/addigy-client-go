package maintenance_entities

import "encoding/json"

type UpdateStagedMaintenanceRequest struct {
	Id             string      `json:"_id"`
	Day            string      `json:"day"`
	Description    string      `json:"description"`
	Enabled        bool        `json:"enabled"`
	Frequency      string      `json:"frequency"`
	Instructions   []string    `json:"instructions"` // Not sure about the type
	IsInBlueprint  bool        `json:"is_in_blueprint"`
	Jobname        string      `json:"jobname"`
	Jobtime        json.Number `json:"jobtime"`
	LocalTime      bool        `json:"local_time"`
	Maxtrycount    int         `json:"maxtrycount"`
	Promptuser     bool        `json:"promptuser"`
	Scheduledtime  string      `json:"scheduledtime"`
	TimeoutSeconds int         `json:"timeout_seconds"`
	Version        int         `json:"version"`
}
