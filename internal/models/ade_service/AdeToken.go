package ade_service

type AdeToken struct {
	Account struct {
		AdminID    string   `json:"admin_id"`
		OrgAddress string   `json:"org_address"`
		OrgEmail   string   `json:"org_email"`
		OrgName    string   `json:"org_name"`
		OrgPhone   string   `json:"org_phone"`
		ServerName string   `json:"server_name"`
		ServerUUID string   `json:"server_uuid"`
		Urls       []string `json:"urls"`
	} `json:"account"`
	DevicesSyncCompleted bool   `json:"devices_sync_completed"`
	Disabled             bool   `json:"disabled"`
	LastScanTime         string `json:"last_scan_time"`
	Orgid                string `json:"orgid"`
	PolicyID             string `json:"policy_id"`
	Removed              bool   `json:"removed"`
	SyncingError         string `json:"syncing_error"`
}
