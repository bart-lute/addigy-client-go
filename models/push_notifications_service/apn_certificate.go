package push_notifications_service

type ApnCertificate struct {
	ApnsLastUpdated string `json:"apns_last_updated"`
	ApnsNotAfter    string `json:"apns_not_after"`
	ApnsNotBefore   string `json:"apns_not_before"`
	ApnsTopic       string `json:"apns_topic"`
	AppleIdUsed     string `json:"apple_id_used"`
	CertIdentifier  string `json:"cert_identifier"`
	CsrId           string `json:"csr_id"`
	IsExpired       bool   `json:"is_expired"`
	UserEmail       string `json:"user_email"`
}
