package mdm_configurations

type PayloadAvailability struct {
	AllowMultiplePayloads   SupportedOs `json:"allow_multiple_payloads"`
	AllowUserEnrollment     SupportedOs `json:"allow_user_enrollment"`
	DeviceChannel           SupportedOs `json:"device_channel"`
	RequiresSupervision     SupportedOs `json:"requires_supervision"`
	RequiresUserApprovedMdm SupportedOs `json:"requires_user_approved_mdm"`
	UserChannel             SupportedOs `json:"user_channel"`
}
