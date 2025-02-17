package instructions_service

type SelfServiceConfiguration struct {
	AppLogo                    File   `json:"app_logo"`
	AssetVersion               int    `json:"asset_version"`
	DarkModeAppLogo            File   `json:"dark_mode_app_logo"`
	Description                string `json:"description"`
	DisplayMode                string `json:"display_mode"`
	DockIcon                   File   `json:"dock_icon"`
	EnrollmentTrigger          string `json:"enrollment_trigger"`
	ExitCode                   string `json:"exit_code"`
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
	InstructionId              string `json:"instruction_id"`
	IntegrationIntuneEnabled   bool   `json:"integration_intune_enabled"`
	IsInBlueprint              bool   `json:"is_in_blueprint"`
	IsOnboardingConfig         bool   `json:"is_onboarding_config"`
	Label                      string `json:"label"`
	MaintenancePromptText      string `json:"maintenance_prompt_text"`
	MenubarIcon                File   `json:"menubar_icon"`
	MsOfficeUpdatesPromptText  string `json:"ms_office_updates_prompt_text"`
	Name                       string `json:"name"`
	OrganizationId             string `json:"organization_id"`
	OsType                     string `json:"os_type"`
	PolicyRestricted           bool   `json:"policy_restricted"`
	Provider                   string `json:"provider"`
	RestartAfterComplete       bool   `json:"restart_after_complete"`
	ScreenviewPromptText       string `json:"screenview_prompt_text"`
	ShowCompleteScreen         bool   `json:"show_complete_screen"`
	ShowDockIcon               bool   `json:"show_dock_icon"`
	ShowInApplications         bool   `json:"show_in_applications"`
	ShowMenubarIcon            bool   `json:"show_menubar_icon"`
	ShowSupport                bool   `json:"show_support"`
	UserSentimentPromptText    string `json:"user_sentiment_prompt_text"`
}
