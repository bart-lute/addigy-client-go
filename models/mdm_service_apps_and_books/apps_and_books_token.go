package mdm_service_apps_and_books

type AppsAndBooksToken struct {
	Active            bool         `json:"active"`
	ClientConfig      ClientConfig `json:"clientConfig"`
	ErrMsg            string       `json:"errMsg"`
	ForceOnPolicyTree bool         `json:"forceOnPolicyTree"`
	InheritedFrom     string       `json:"inheritedFrom"`
	LastSync          string       `json:"lastSync"`
	LastSyncingStatus string       `json:"lastSyncingStatus"`
	LocationId        string       `json:"locationId"`
	OrgName           string       `json:"orgName"`
	PolicyId          string       `json:"policyId"`
	Syncing           bool         `json:"syncing"`
	Version           string       `json:"version"`
}
