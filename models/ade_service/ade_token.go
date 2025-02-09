package ade_service

type AdeToken struct {
	Account              AdeAccount `json:"account"`
	DevicesSyncCompleted bool       `json:"devices_sync_completed"`
	Disabled             bool       `json:"disabled"`
	LastScanTime         string     `json:"last_scan_time"`
	Orgid                string     `json:"orgid"`
	PolicyId             string     `json:"policy_id"`
	Removed              bool       `json:"removed"`
	SyncingError         string     `json:"syncing_error"`
}
