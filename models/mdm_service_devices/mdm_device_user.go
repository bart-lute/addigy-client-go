package mdm_service_devices

type MdmDeviceUser struct {
	DataQuota     int    `json:"data_quota"`
	DataUsed      int    `json:"data_used"`
	FullName      string `json:"full_name"`
	HasDataToSync bool   `json:"has_data_to_sync"`
	MobileAccount bool   `json:"mobile_account"`
	Orgid         string `json:"orgid"`
	Udid          string `json:"udid"`
	UserName      string `json:"user_name"`
	UserUid       int    `json:"user_uid"`
}
