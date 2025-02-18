package mdm_configurations

type SupportedOsVersions struct {
	Os       string            `json:"os"`
	Versions SupportedVersions `json:"versions"`
}
