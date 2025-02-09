package response_entities

type Response struct {
	Items    map[string]string `json:"items"`
	Metadata Metadata          `json:"metadata"`
}
