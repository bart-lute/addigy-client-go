package default_assets_service

type DefaultSelfServiceConfiguration struct {
	Condition                  string `json:"condition"`
	FilevaultPromptText        string `json:"filevault_prompt_text"`
	HideChat                   bool   `json:"hide_chat"`
	HomeScreenAddress          string `json:"home_screen_address"`
	HomeScreenCompanyName      string `json:"home_screen_company_name"`
	HomeScreenConfigureDetails bool   `json:"home_screen_configure_details"`
	HomeScreenDescription      string `json:"home_screen_description"`
	HomeScreenEmail            string `json:"home_screen_email"`
	HomeScreenPhone            string `json:"home_screen_phone"`
	HomeScreenShowAddress      bool   `json:"home_screen_show_address"`
	HomeScreenShowDescription  bool   `json:"home_screen_show_description"`
	HomeScreenShowEmail        bool   `json:"home_screen_show_email"`
	HomeScreenShowPhone        bool   `json:"home_screen_show_phone"`
	Id                         string `json:"id"`
	InstructionId              string `json:"instruction_id"`
	IntegrationIntuneEnabled   bool   `json:"integration_intune_enabled"`
	Label                      string `json:"label"`
	MaintenancePromptText      string `json:"maintenance_prompt_text"`
	MsOfficeUpdatesPromptText  string `json:"ms_office_updates_prompt_text"`
	Name                       string `json:"name"`
	PolicyRestricted           bool   `json:"policy_restricted"`
	Provider                   string `json:"provider"`
	RemoveScript               string `json:"remove_script"`
	RunOnSuccess               bool   `json:"run_on_success"`
	ScreenviewPromptText       string `json:"screenview_prompt_text"`
	ShortDescription           string `json:"short_description"`
	ShowDockIcon               bool   `json:"show_dock_icon"`
	ShowInApplications         bool   `json:"show_in_applications"`
	ShowMenubarIcon            bool   `json:"show_menubar_icon"`
	ShowSupport                bool   `json:"show_support"`
	StatusOnSkipped            string `json:"status_on_skipped"`
	UserSentimentPromptText    string `json:"user_sentiment_prompt_text"`
}
