package device_entities

type DeviceFact struct {
	ErrorMsg string `json:"error_msg"`
	Type     string `json:"type"`
	Value    any    `json:"value"`
}
