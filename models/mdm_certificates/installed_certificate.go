package mdm_certificates

type InstalledCertificate struct {
	CertOrgName      string `json:"cert_org_name"`
	CommonName       string `json:"common_name"`
	Data             []byte `json:"data"` // integer($binary) ???
	DeviceUuid       string `json:"device_uuid"`
	IsIdentity       bool   `json:"is_identity"`
	IssuerCommonName string `json:"issuer_common_name"`
	NotAfter         string `json:"not_after"`
	NotBefore        string `json:"not_before"`
	UserId           string `json:"user_id"`
	Version          int    `json:"version"`
}
