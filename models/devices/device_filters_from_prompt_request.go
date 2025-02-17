package devices

import "encoding/json"

type DeviceFiltersFromPromptRequest struct {
	FactCount      int         `json:"fact_count"`
	PromptText     string      `json:"prompt_text"`
	ResponseLength int         `json:"response_length"`
	Temperature    json.Number `json:"temperature"`
}
