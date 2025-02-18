package mdm_configurations

type ConfigurationPayloads struct {
	Payloads            []interface{} `json:"payloads"`
	PoliciesMdmPayloads []interface{} `json:"policies_mdm_payloads"`
	StagedPayloads      []interface{} `json:"staged_payloads"`
}
