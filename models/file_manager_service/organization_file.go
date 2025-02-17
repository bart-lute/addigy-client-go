package file_manager_service

type OrganizationFile struct {
	ContentType string `json:"content_type"`
	Created     string `json:"created"`
	Filename    string `json:"filename"`
	Id          string `json:"id"`
	Md5Hash     string `json:"md5_hash"`
	Orgid       string `json:"orgid"`
	Provider    string `json:"provider"`
	Size        int    `json:"size"`
	UserEmail   string `json:"user_email"`
}
