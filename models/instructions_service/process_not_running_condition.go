package instructions_service

type ProcessNotRunningCondition struct {
	Enabled     bool   `json:"enabled"`
	ProcessName string `json:"process_name"`
}
