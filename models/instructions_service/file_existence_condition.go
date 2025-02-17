package instructions_service

type FileExistenceCondition struct {
	Enabled bool   `json:"enabled"`
	Path    string `json:"path"`
}
