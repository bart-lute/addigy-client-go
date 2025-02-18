package mdm_commands

type RestartCommandExample struct {
	KextPaths                    []string `json:"KextPaths"`
	RebuildKernelCache           bool     `json:"RebuildKernelCache"`
	RequestRequiresNetworkTether bool     `json:"RequestRequiresNetworkTether"`
	RequestType                  string   `json:"RequestType"`
}
