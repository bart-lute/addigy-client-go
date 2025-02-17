package events_service

import (
	"github.com/bart-lute/addigy-client-go/models/events_client"
	"github.com/bart-lute/addigy-client-go/models/response_entities"
)

type ListEventsPublicResponse struct {
	Items       []events_client.EventResponseItem `json:"items"`
	Keywords    []string                          `json:"keywords"`
	Metadata    response_entities.Metadata        `json:"metadata"`
	SearchAfter int                               `json:"search_after"`
	Took        int                               `json:"took"`
}
