package apps_and_books_service

type Configuration struct {
	Name  string            `json:"name"`
	Type  string            `json:"type"`
	Value map[string]string `json:"value"` // Revisit later
}
