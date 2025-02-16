package device_script_assignments

import "github.com/bart-lute/addigy-client-go/models/scripts_service"

type GetDeviceScriptAssignmentsResponse struct {
	Items []scripts_service.DeviceScriptAssignmentResponse `json:"items"`
}
