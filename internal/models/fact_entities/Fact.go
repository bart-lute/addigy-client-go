package fact_entities

type Fact struct {
	CommunityFactId  string          `json:"community_fact_id"`
	CommunityVersion int             `json:"community_version"`
	Id               string          `json:"id"`
	Name             string          `json:"name"`
	Notes            string          `json:"notes"`
	OsArchitectures  OSArchitectures `json:"os_architectures"`
	ReturnType       string          `json:"return_type"`
	Version          int             `json:"version"`
}
