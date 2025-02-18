package mbov_service

type AccountMetadata struct {
	AccountCreationCompleted bool   `json:"account_creation_completed"`
	AccountId                string `json:"account_id"`
	CreatedDate              string `json:"created_date"`
	Email                    string `json:"email"`
}
