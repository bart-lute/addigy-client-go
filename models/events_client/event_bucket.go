package events_client

type EventBucket struct {
	Buckets []StatusBucket `json:"buckets"`
	Count   int            `json:"count"`
	Date    string         `json:"date"`
}
