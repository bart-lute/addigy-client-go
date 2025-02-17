package feature_betas_entities

type FeatureBeta struct {
	Description    string `json:"description"`
	FeatureFlagKey string `json:"feature_flag_key"`
	HtmlDetails    string `json:"html_details"`
	KbLink         string `json:"kb_link"`
	Name           string `json:"name"`
	Public         bool   `json:"public"`
}
