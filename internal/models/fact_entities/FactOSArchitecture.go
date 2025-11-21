package fact_entities

type FactOSArchitecture struct {
	IsSupported bool   `json:"is_supported"`
	Language    string `json:"language"`
	Script      string `json:"script"`
	Shebang     string `json:"shebang"`
}
