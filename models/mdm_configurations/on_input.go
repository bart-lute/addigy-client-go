package mdm_configurations

type OnInput struct {
	Action            string            `json:"action"`
	Condition         OnSelectCondition `json:"condition"`
	TargetKeyVariable string            `json:"target_key_variable"`
	Value             string            `json:"value"` // ???
}
