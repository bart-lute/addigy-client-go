package instructions_service

type ProfileEvent struct {
	Application Application `json:"application"`
	Permission  string      `json:"permission"`
}
