package feature_betas_entities

import "github.com/bart-lute/addigy-client-go/models/response_entities"

type FeatureBetaResponse struct {
	Items    []FeatureBeta              `json:"items"`
	Metadata response_entities.Metadata `json:"metadata"`
}
