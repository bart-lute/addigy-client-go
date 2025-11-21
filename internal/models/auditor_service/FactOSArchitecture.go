package auditor_service

type FactOSArchitecture struct {
	IsSupported bool   `json:"is_supported"`
	Language    string `json:"language"`
	Md5Hash     string `json:"md5_hash"`
	Script      string `json:"script"`
	Shebang     string `json:"shebang"`
}
