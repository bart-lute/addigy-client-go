package mbov_service

type Site struct {
	CompanyName     string `json:"company_name"`
	Email           string `json:"email"`
	Firstname       string `json:"firstname"`
	Id              string `json:"id"`
	Lastname        string `json:"lastname"`
	NebulaAccountId string `json:"nebula_account_id"`
	SiteEndDate     string `json:"site_end_date"`
}
