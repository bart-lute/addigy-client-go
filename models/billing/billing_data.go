package billing

import "github.com/bart-lute/addigy-client-go/models/billing_service"

type BillingData struct {
	AccountSuspended  bool                                      `json:"account_suspended"`
	BillingAllowed    bool                                      `json:"billing_allowed"`
	DevicesCount      []billing_service.OrganizationDeviceCount `json:"devices_count"`
	InTrial           bool                                      `json:"in_trial"`
	ParentCompanyName string                                    `json:"parent_company_name"`
	Subscriptions     []billing_service.Subscription            `json:"subscriptions"`
	TrialEndDate      string                                    `json:"trial_end_date"`
}
