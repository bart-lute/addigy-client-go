package mdm_service_devices

import "encoding/json"

type LocationInfo struct {
	Altitude           json.Number `json:"altitude"`
	Course             json.Number `json:"course"`
	HorizontalAccuracy json.Number `json:"horizontal_accuracy"`
	Latitude           json.Number `json:"latitude"`
	Longitude          json.Number `json:"longitude"`
	Speed              json.Number `json:"speed"`
	Timestamp          string      `json:"timestamp"`
	VerticalAccuracy   json.Number `json:"vertical_accuracy"`
}
