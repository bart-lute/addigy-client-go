package policy_service

type Image struct {
	ContentType    string `json:"content_type"`
	Created        string `json:"created"`
	FileName       string `json:"file_name"`
	Id             string `json:"id"`
	Md5Hash        string `json:"md_5_hash"`
	OrganizationId string `json:"organization_id"`
	Provider       string `json:"provider"`
	Size           int    `json:"size"`
	UserEmail      string `json:"user_email"`
}
