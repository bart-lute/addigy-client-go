package device_users_service

type Error struct {
	Message   string `json:"message"`
	RowNumber int    `json:"row_number"`
}
