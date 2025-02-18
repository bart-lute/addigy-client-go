package mdm_devices

import (
	"github.com/bart-lute/addigy-client-go/models/mdm_service_devices"
	"github.com/bart-lute/addigy-client-go/models/mdm_service_profiles"
	"github.com/bart-lute/addigy-client-go/models/push_notifications_service"
)

type MdmDetailsResponse struct {
	ApnCertificate    push_notifications_service.ApnCertificate    `json:"apn_certificate"`
	EnrollmentProfile mdm_service_profiles.DeviceEnrollmentProfile `json:"enrollment_profile"`
	LastResponse      mdm_service_devices.CommandResponseStatus    `json:"last_response"`
}
