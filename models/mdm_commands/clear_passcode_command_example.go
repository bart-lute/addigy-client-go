package mdm_commands

type ClearPasscodeCommandExample struct {
	RequestRequiresNetworkTether bool   `json:"RequestRequiresNetworkTether"`
	RequestType                  string `json:"RequestType"`
	UnlockToken                  string `json:"UnlockToken"`
}
