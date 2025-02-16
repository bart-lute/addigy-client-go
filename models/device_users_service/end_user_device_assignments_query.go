package device_users_service

type EndUserDeviceAssignmentsQuery struct {
	EndUserIds    []string `json:"end_user_ids"`
	SerialNumbers []string `json:"serial_numbers"`
}
