package billing_service

type CardUpdate struct {
	AddressCity    string `json:"address_city"`
	AddressCountry string `json:"address_country"`
	AddressLine1   string `json:"address_line_1"`
	AddressLine2   string `json:"address_line_2"`
	AddressState   string `json:"address_state"`
	AddressZip     string `json:"address_zip"`
	CardId         string `json:"card_id"`
	ExpMonth       int    `json:"exp_month"`
	ExpYear        int    `json:"exp_year"`
	Name           string `json:"name"`
}
