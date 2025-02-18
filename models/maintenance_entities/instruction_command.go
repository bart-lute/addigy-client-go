package maintenance_entities

type InstructionCommand struct {
	Name   string `json:"name"`
	Script string `json:"script"`
	Type   string `json:"type"`
}
