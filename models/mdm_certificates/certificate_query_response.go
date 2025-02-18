package mdm_certificates

import "github.com/bart-lute/addigy-client-go/models/response_entities"

type CertificateQueryResponse struct {
	Items    []InstalledCertificate     `json:"items"`
	Metadata response_entities.Metadata `json:"metadata"`
}
