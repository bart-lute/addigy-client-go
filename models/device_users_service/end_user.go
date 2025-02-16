package device_users_service

type EndUser struct {
	Department     string `json:"department"`
	Email          string `json:"email"`
	FullName       string `json:"full_name"`
	Id             string `json:"id"`
	Location       string `json:"location"`
	OrganizationId string `json:"organization_id"`
	Position       string `json:"position"`
	Username       string `json:"username"`
}
