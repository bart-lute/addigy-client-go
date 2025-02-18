package mdm_profiles

import "github.com/bart-lute/addigy-client-go/models/mdm_configurations"

type MDMProfileListResponse struct {
	Count   int                              `json:"count"`
	Results []mdm_configurations.MDMProfiles `json:"results"`
	total   int                              `json:"total"`
}
