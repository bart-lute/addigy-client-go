package events_client

type EventResponseItem struct {
	Action         Action            `json:"action"`
	ActionReceiver Actor             `json:"action_receiver"`
	ActionSender   Actor             `json:"action_sender"`
	Date           string            `json:"date"`
	EventId        string            `json:"event_id"`
	Highlights     map[string]string `json:"highlights"`
	Level          string            `json:"level"`
	Orgid          string            `json:"orgid"`
	Result         Result            `json:"result"`
	Source         string            `json:"source"`
}
