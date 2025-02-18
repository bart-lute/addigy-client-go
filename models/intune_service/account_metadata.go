package intune_service

type AccountMetadata struct {
	AzurePartnerComplianceGroupId string `json:"azure_partner_compliance_group_id"`
	CreationDate                  string `json:"creation_date"`
	CustomerEnrollmentUrl         string `json:"customer_enrollment_url"`
	CustomerTenantId              string `json:"customer_tenant_id"`
	Enabled                       bool   `json:"enabled"`
	IsCompliantByDefault          bool   `json:"is_compliant_by_default"`
	LastProvisionDate             string `json:"last_provision_date"`
	OrganizationId                string `json:"organization_id"`
	PolicyId                      string `json:"policy_id"`
}
