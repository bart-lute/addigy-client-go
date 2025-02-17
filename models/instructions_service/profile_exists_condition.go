package instructions_service

type ProfileExistsCondition struct {
	Enabled   bool   `json:"enabled"`
	ProfileId string `json:"profile_id"`
}
