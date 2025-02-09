package alert_entities

type AlertsScheduleRequest struct {
	ExcludedIds   []string `json:"excluded_ids"`
	Ids           []string `json:"ids"`
	Limit         int      `json:"limit"`
	NameContains  string   `json:"name_contains"`
	Skip          int      `json:"skip"`
	SortDirection string   `json:"sort_direction"` // asc, desc
	SortField     string   `json:"sort_field"`
}
