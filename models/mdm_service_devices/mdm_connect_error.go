package mdm_service_devices

type MdmConnectError struct {
	ErrorCode            int    `json:"error_code"`
	ErrorDomain          string `json:"error_domain"`
	LocalizedDescription string `json:"localized_description"`
	UsEnglishDescription string `json:"us_english_description"`
}
