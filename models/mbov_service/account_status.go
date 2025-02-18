package mbov_service

type AccountStatus struct {
	Account        AccountMetadata `json:"account"`
	Enabled        bool            `json:"enabled"`
	MbovActive     bool            `json:"mbov_active"`
	SuspendPending bool            `json:"suspend_pending"`
}
