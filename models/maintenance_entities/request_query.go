package maintenance_entities

type RequestQuery struct {
	Page          int    `json:"page"`
	PerPage       int    `json:"per_page"`
	Query         Filter `json:"query"`
	SortDirection string `json:"sort_direction"`
	SortField     string `json:"sort_field"`
}
