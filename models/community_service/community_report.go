package community_service

type CommunityReport struct {
	Id           string `json:"id"`
	ReportReason string `json:"report_reason"`
	Type         string `json:"type"` // Enum: [ custom-fact, predefined-command ]
}
