package installed_applications_service

type InstalledApplication struct {
	Agentid                 string `json:"agentid"`
	MicrosoftBuildIncrement string `json:"microsoft_build_increment"`
	MicrosoftBuildNumber    string `json:"microsoft_build_number"`
	Name                    string `json:"name"`
	Orgid                   string `json:"orgid"`
	Path                    string `json:"path"`
	Version                 string `json:"version"`
}
