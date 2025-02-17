package dm_checkin_service

type PendingUpdate struct {
	BuildVersion        string   `json:"build_version"`
	OsVersion           string   `json:"os_version"`
	Reasons             []string `json:"reasons"`
	Status              string   `json:"status"`
	TargetLocalDateTime string   `json:"target_local_date_time"`
}
