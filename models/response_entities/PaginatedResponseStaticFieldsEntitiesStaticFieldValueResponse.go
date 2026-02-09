package response_entities

type StaticFieldsValueItem struct {
	Agentid       string `json:"agentid"`
	StaticFieldID string `json:"static_field_id"`
	Value         string `json:"value"`
}

type PaginatedResponseStaticFieldsEntitiesStaticFieldValueResponse struct {
	Items    []StaticFieldsValueItem `json:"items"`
	Metadata struct {
		Page        int `json:"page"`
		PageCount   int `json:"page_count"`
		PerPage     int `json:"per_page"`
		ResultCount int `json:"result_count"`
		Total       int `json:"total"`
	} `json:"metadata"`
}
