package fact_entities

type FactPutRequest struct {
	Id              string          `json:"id"`
	Name            string          `json:"name"`
	Notes           string          `json:"notes"`
	OsArchitectures OSArchitectures `json:"os_architectures"`
	ReturnType      string          `json:"return_type"`
}
