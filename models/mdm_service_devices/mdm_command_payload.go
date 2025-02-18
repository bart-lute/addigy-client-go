package mdm_service_devices

type MDMCommandPayload struct {
	Command           interface{}       `json:"Command"`
	CommandUUID       string            `json:"CommandUUID"`
	CreatedDate       string            `json:"CreatedDate"`
	DeviceUDID        string            `json:"DeviceUDID"`
	ErrorChain        []MdmConnectError `json:"ErrorChain"`
	ExecutionDeadline string            `json:"ExecutionDeadline"`
	ExpirationDate    string            `json:"ExpirationDate"`
	LastUpdated       string            `json:"LastUpdated"`
	ManagedUserID     string            `json:"ManagedUserID"`
	Name              string            `json:"Name"`
	StartingDeadline  string            `json:"StartingDeadline"`
	Status            string            `json:"Status"`
	Weight            int               `json:"Weight"`
}
