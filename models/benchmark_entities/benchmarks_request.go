package benchmark_entities

type BenchmarksRequest struct {
	Page          int    `json:"page"`
	PerPage       int    `json:"per_page"`
	Query         Filter `json:"query"`
	SortDirection string
	SortField     string
}
