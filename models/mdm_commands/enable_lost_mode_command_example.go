package mdm_commands

type EnableLostModeCommandExample struct {
	Footnote                     string `json:"Footnote"`
	Message                      string `json:"Message"`
	PhoneNumber                  string `json:"PhoneNumber"`
	RequestRequiresNetworkTether bool   `json:"RequestRequiresNetworkTether"`
	RequestType                  string `json:"RequestType"`
}
