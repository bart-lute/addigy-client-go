package instructions_service

type AppExistsCondition struct {
	Enabled             bool   `json:"enabled"`
	InstallIfNotPresent bool   `json:"install_if_not_present"`
	Operator            string `json:"operator"` // [ eq, lt, gt, lte, gte ]
	Path                string `json:"path"`
	Version             string `json:"version"`
}
