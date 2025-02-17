package facts

type FactUsage struct {
	Alerts       []UsageItem `json:"alerts"`
	Integrations []UsageItem `json:"integrations"`
	Policies     []UsageItem `json:"policies"`
	Reports      []UsageItem `json:"reports"`
	UserConfigs  []UsageItem `json:"user_configs"`
}
