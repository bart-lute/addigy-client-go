package mdm_service_devices

type EraseDeviceCommandExample struct {
	DisallowProximitySetup bool   `json:"DisallowProximitySetup"`
	ObliterationBehavior   string `json:"ObliterationBehavior"`
	PIN                    string `json:"PIN"`
	PreserveDataPlan       bool   `json:"PreserveDataPlan"`
}
