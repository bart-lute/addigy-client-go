package events_client

type Query struct {
	Fields []string `json:"fields"` // Example: ["action_sender.type"]
	Query  string   `json:"query"`
}
