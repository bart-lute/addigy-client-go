package end_users

type PaginatedEndUsersImportMetadataQuery struct {
	Page          int                         `json:"page"`
	PerPage       int                         `json:"per_page"`
	Query         EndUsersImportMetadataQuery `json:"query"`
	SortDirection string                      `json:"sort_direction"`
	SortField     string                      `json:"sort_field"`
}
