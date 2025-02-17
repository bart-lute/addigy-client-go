package events_client

type SearchEventOptions struct {
	AggregationIntervalMinutes int  `json:"aggregation_interval_minutes"`
	Highlight                  bool `json:"highlight"`
	Keywords                   bool `json:"keywords"`
}
