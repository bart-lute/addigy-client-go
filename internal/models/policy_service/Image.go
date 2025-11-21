package policy_service

type Image struct {
	ContentType    string `json:"content_type"`
	Created        string `json:"created"`
	FileName       string `json:"file_name"`
	ID             string `json:"id"`
	Md5Hash        string `json:"md_5_hash"`
	OrganizationID string `json:"organization_id"`
	Provider       string `json:"provider"`
	Size           int    `json:"size"`
	UserEmail      string `json:"user_email"`
}
