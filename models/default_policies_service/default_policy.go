package default_policies_service

type DefaultPolicy struct {
	Color                              string        `json:"color"`
	DefaultComplianceAssignments       []Assignments `json:"default_compliance_assignments"`
	DefaultMdmConfigurationAssignments []Assignments `json:"default_mdm_configuration_assignments"`
	DefaultMonitoringAssignments       []Assignments `json:"default_monitoring_assignments"`
	DefaultSelfServiceAssignments      []Assignments `json:"default_self_service_assignments"`
	DefaultSoftwareAssignments         []Assignments `json:"default_software_assignments"`
	Description                        string        `json:"description"`
	Icon                               string        `json:"icon"`
	Id                                 string        `json:"id"`
	Identifier                         string        `json:"identifier"`
	Name                               string        `json:"name"`
	Order                              int           `json:"order"`
	SortDescription                    string        `json:"sort_description"`
	Version                            int           `json:"version"`
}
