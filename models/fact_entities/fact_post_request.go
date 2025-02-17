package fact_entities

type FactPostRequest struct {
	Name            string          `json:"name"`
	Notes           string          `json:"notes"`
	OsArchitectures OSArchitectures `json:"os_architectures"`
	ReturnType      string          `json:"return_type"`
}
