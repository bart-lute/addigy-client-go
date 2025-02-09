package alerts

type AlertPolicyRequest struct {
	AlertId  string `json:"alert_id"`
	PolicyId string `json:"policy_id"`
}
