package billing_service

type Card struct {
	AddressCity    string `json:"address_city"`
	AddressCountry string `json:"address_country"`
	AddressLine1   string `json:"address_line_1"`
	AddressLine2   string `json:"address_line_2"`
	AddressState   string `json:"address_state"`
	AddressZip     string `json:"address_zip"`
	Brand          string `json:"brand"`
	ExpMonth       int    `json:"exp_month"`
	ExpYear        int    `json:"exp_year"`
	Id             string `json:"id"`
	IsDefault      bool   `json:"is_default"`
	Last4          string `json:"last_4"`
	Name           string `json:"name"`
}
