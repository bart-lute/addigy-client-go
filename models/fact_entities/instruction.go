package fact_entities

type Instruction struct {
	Condition        string `json:"condition"`
	Identifier       string `json:"identifier"`
	InstructionId    string `json:"instruction_id"`
	Label            string `json:"label"`
	Name             string `json:"name"`
	PolicyRestricted bool   `json:"policy_restricted"`
	Provider         string `json:"provider"`
	Public           bool   `json:"public"`
	RemoveScript     string `json:"remove_script"`
	RunOnSuccess     bool   `json:"run_on_success"`
	StatusOnSkipped  string `json:"status_on_skipped"`
	UserEmail        string `json:"user_email"`
}
