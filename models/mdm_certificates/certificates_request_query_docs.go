package mdm_certificates

type CertificatesRequestQueryDocs struct {
	DaysUntilExpiration int      `json:"days_until_expiration"`
	Devices             []string `json:"devices"`
	IssuerCommonName    string   `json:"issuer_common_name"`
}
