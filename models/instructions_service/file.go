package instructions_service

type File struct {
	ContentType string `json:"content_type"`
	Created     string `json:"created"`
	FilePath    string `json:"file_path"`
	Filename    string `json:"filename"`
	Id          string `json:"id"`
	Md5Hash     string `json:"md5_hash"`
	Orgid       string `json:"orgid"`
	Provider    string `json:"provider"`
	Size        int    `json:"size"`
	UserEmail   string `json:"user_email"`
}
