package mdm_service_apps_and_books

type ClientConfig struct {
	CountryISO2ACode            string              `json:"countryISO2ACode"`
	DefaultPlatform             string              `json:"defaultPlatform"`
	LocationName                string              `json:"locationName"`
	MdmInfo                     ClientConfigMdmInfo `json:"mdmInfo"`
	SubscribedNotificationTypes []string            `json:"subscribedNotificationTypes"`
	UId                         string              `json:"uId"`
	WebsiteURL                  string              `json:"websiteURL"`
}
