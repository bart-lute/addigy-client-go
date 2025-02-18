package mdm_configurations

type UI struct {
	Columns         []Column  `json:"columns"`
	Component       string    `json:"component"`
	Label           string    `json:"label"`
	Message         string    `json:"message"`
	OnInput         []OnInput `json:"on_input"`
	Placeholder     string    `json:"placeholder"`
	Required        bool      `json:"required"`
	SupportedValues []string  `json:"supported_values"` // ???
	Tags            []string  `json:"tags"`
	Type            string    `json:"type"`
	Variable        string    `json:"variable"`
}
