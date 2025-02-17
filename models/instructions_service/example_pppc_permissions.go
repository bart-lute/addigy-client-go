package instructions_service

type ExamplePPPCPermissions struct {
	Accessibility                string `json:"Accessibility"`
	AddressBook                  string `json:"AddressBook"`
	Calendar                     string `json:"Calendar"`
	Camera                       string `json:"Camera"`
	FileProviderPresence         string `json:"FileProviderPresence"`
	ListenEvent                  string `json:"ListenEvent"`
	MediaLibrary                 string `json:"MediaLibrary"`
	Microphone                   string `json:"Microphone"`
	Photos                       string `json:"Photos"`
	PostEvent                    string `json:"PostEvent"`
	Reminders                    string `json:"Reminders"`
	ScreenCapture                string `json:"ScreenCapture"`
	SpeechRecognition            string `json:"SpeechRecognition"`
	SystemPolicyAllFiles         string `json:"SystemPolicyAllFiles"`
	SystemPolicyDesktopFolder    string `json:"SystemPolicyDesktopFolder"`
	SystemPolicyDocumentsFolder  string `json:"SystemPolicyDocumentsFolder"`
	SystemPolicyDownloadsFolder  string `json:"SystemPolicyDownloadsFolder"`
	SystemPolicyNetworkVolumes   string `json:"SystemPolicyNetworkVolumes"`
	SystemPolicyRemovableVolumes string `json:"SystemPolicyRemovableVolumes"`
	SystemPolicySysAdminFiles    string `json:"SystemPolicySysAdminFiles"`
}
