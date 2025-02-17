package end_users

import "github.com/bart-lute/addigy-client-go/models/response_entities"

type EndUsersQueryResponse struct {
	Items    []EndUsersQueryResponse    `json:"items"`
	Metadata response_entities.Metadata `json:"metadata"`
}
