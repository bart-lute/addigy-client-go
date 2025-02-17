package fact_entities

type Response struct {
	Failed    []string `json:"failed"`
	Succeeded []string `json:"succeeded"`
}
