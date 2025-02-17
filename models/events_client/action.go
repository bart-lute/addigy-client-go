package events_client

type Action struct {
	Details string `json:"details"`
	Entity  Actor  `json:"entity"`
	Name    string `json:"name"`
}
