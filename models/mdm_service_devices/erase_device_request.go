package mdm_service_devices

type EraseDeviceRequest struct {
	DeviceUuid             string `json:"device_uuid"`
	DisallowProximitySetup bool   `json:"disallow_proximity_setup"`
	ObliterationBehavior   string `json:"obliteration_behavior"`
	Pin                    string `json:"pin"`
	PreserveDataPlan       bool   `json:"preserve_data_plan"`
}
