package addigy

type QueryRequest struct {
	DesiredFactIdentifiers []string `json:"desired_fact_identifiers,omitempty"`
	Page                   int      `json:"page"`
	PerPage                int      `json:"per_page"`
	SortDirection          string   `json:"sort_direction"`
	SortField              string   `json:"sort_field"`
	Query                  any      `json:"query"`
}

type SmartSoftwareQuery struct {
	Archived     bool     `json:"archived,omitempty"`
	ExcludedIds  []string `json:"excluded_ids,omitempty"`
	Identifier   string   `json:"identifier,omitempty"`
	Ids          []string `json:"ids,omitempty"`
	NameContains string   `json:"name_contains,omitempty"`
}

type CustomFactsQuery struct {
	Id           string `json:"id,omitempty"`
	NameContains string `json:"name_contains,omitempty"`
}

type VariablesQuery struct {
	KeyContains string   `json:"key_contains"`
	Keys        []string `json:"keys"`
}

// MetaData Returned as part of a Response
type MetaData struct {
	Page        int `json:"page"`
	PageCount   int `json:"page_count"`
	PerPage     int `json:"per_page"`
	ResultCount int `json:"result_count"`
	Total       int `json:"total"`
}

// QueryResponse Response Data when paged data is expected
type QueryResponse struct {
	Items    []any    `json:"items"`
	Metadata MetaData `json:"metadata"`
}

type SmartSoftwareItem struct {
	Archived                    bool    `json:"archived"`
	BaseIdentifier              string  `json:"base_identifier"`
	Category                    string  `json:"category"`
	Description                 string  `json:"description"`
	FactIdentifier              string  `json:"fact_identifier"`
	Identifier                  string  `json:"identifier"`
	Install                     bool    `json:"install"`
	InstallationScript          string  `json:"installation_script"`
	InstructionId               string  `json:"instruction_id"`
	PolicyRestricted            bool    `json:"policy_restricted"`
	PricePerDevice              float64 `json:"price_per_device"`
	Priority                    float64 `json:"priority"`
	Provider                    string  `json:"provider"`
	Public                      bool    `json:"public"`
	PublicSoftwareInstructionId string  `json:"public_software_instruction_id"`
	RemoveScript                string  `json:"remove_script"`
	RunOnSuccess                bool    `json:"run_on_success"`
	SetGroupName                bool    `json:"set_group_name"`
	StatusOnSkipped             string  `json:"status_on_skipped"`
	TccVersion                  int     `json:"tcc_version"`
	Type                        string  `json:"type"`
	UserEmail                   string  `json:"user_email"`
	Version                     string  `json:"version"`
	//Downloads            []string `json:"downloads"` // todo
	//PredefinedConditions []string `json:"predefined_conditions"` // todo
	//Profiles	[]string	// todo
	//SoftwareIcon	// todo
}

type CustomFactsItem struct {
	CommunityFactId  string `json:"community_fact_id"`
	CommunityVersion int    `json:"community_version"`
	Id               string `json:"id"`
	Name             string `json:"name"`
	Notes            string `json:"notes"`
	//OsArchitectures []string `json:"os_architectures"` // todo
	ReturnType string `json:"return_type"`
	Version    int    `json:"version"`
}

type VariablesItem struct {
	CreatedDate  string `json:"created_date"`
	DefaultValue string `json:"default_value"`
	Key          string `json:"key"`
	Type         string `json:"type"`
	UpdatedDate  string `json:"updated_date"`
}
