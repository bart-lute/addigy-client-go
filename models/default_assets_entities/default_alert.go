package default_assets_entities

type DefaultAlert struct {
	Category           string        `json:"category"`
	Enabled            bool          `json:"enabled"`
	Fact               string        `json:"fact"`
	FactIdentifier     string        `json:"fact_identifier"`
	Id                 string        `json:"id"`
	Instructions       []interface{} `json:"instructions"`
	Level              string        `json:"level"`
	Name               string        `json:"name"`
	RemediationEnabled bool          `json:"remediation_enabled"`
	RemediationTime    int           `json:"remediation_time"`
	Selector           string        `json:"selector"`
	ShortDescription   string        `json:"short_description"`
	Type               string        `json:"type"`
	Value              interface{}   `json:"value"`
	ValueType          string        `json:"value_type"`
}
