package benchmarks

import (
	"github.com/bart-lute/addigy-client-go/models/compliance_service"
	"github.com/bart-lute/addigy-client-go/models/response_entities"
)

type BenchmarksQueryResponse struct {
	Items    []compliance_service.Benchmark `json:"items"`
	Metadata response_entities.Metadata     `json:"metadata"`
}
