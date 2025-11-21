package facts

type Filter struct {
	Ids          []string `json:"ids"`
	NameContains string   `json:"name_contains"`
}
