package instructions_service

type OsVersionCondition struct {
	Enabled  bool   `json:"enabled"`
	Operator string `json:"operator"` // [ eq, lt, gt, lte, gte ]
	Version  string `json:"version"`
}
