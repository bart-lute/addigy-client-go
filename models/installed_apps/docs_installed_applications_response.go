package installed_apps

import "github.com/bart-lute/addigy-client-go/models/response_entities"

type DocsInstalledApplicationsResponse struct {
	Items    []InstalledApplication     `json:"items"`
	Metadata response_entities.Metadata `json:"metadata"`
}
