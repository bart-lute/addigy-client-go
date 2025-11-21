package facts

type PaginatedCustomFactsRequestQuery struct {
	Page          int     `json:"page"`
	PerPage       int     `json:"per_page"`
	Query         *Filter `json:"query"`
	SortDirection string  `json:"sort_direction"` // asc | desc
	SortField     string  `json:"sort_field"`
}
