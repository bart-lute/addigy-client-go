package policy_service

type PrebuiltAppSettings struct {
	AdminUpdateDeferralDays int `json:"admin_update_deferral_days"`
	PromptIntervalHours     int `json:"prompt_interval_hours"`
}
