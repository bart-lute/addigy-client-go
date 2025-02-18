package mdm_certificates

type CertificatesRequestDocs struct {
	Page    int                          `json:"page"`
	PerPage int                          `json:"per_page"`
	Query   CertificatesRequestQueryDocs `json:"query"`
}
