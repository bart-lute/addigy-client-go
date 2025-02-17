package files

type OrganizationFilesRequest struct {
	Md5Hash       []string `json:"md5_hash"`
	Page          int      `json:"page"`
	PerPage       int      `json:"per_page"`
	SearchTerm    string   `json:"search_term"`
	SortDirection string   `json:"sort_direction"`
	SortField     string   `json:"sort_field"`
}
