package device_users_service

type EndUsersImportMetadata struct {
	Errors             []Error `json:"errors"`
	FileId             string  `json:"file_id"`
	Filename           string  `json:"filename"`
	Name               string  `json:"name"`
	OrganizationId     string  `json:"organization_id"`
	ProcessedTimestamp string  `json:"processed_timestamp"`
	RequestedTimestamp string  `json:"requested_timestamp"`
	RowCount           int     `json:"row_count"`
	Status             string  `json:"status"`
}
