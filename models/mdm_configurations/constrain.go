package mdm_configurations

type Constrain struct {
	ConstrainsKeys  []ConstrainKey `json:"constrains_keys"`
	LogicalOperator string         `json:"logical_operator"`
	Message         string         `json:"message"`
	Value           interface{}    `json:"value"`
}
