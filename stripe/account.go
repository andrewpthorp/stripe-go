package stripe

type Account struct {
	Id                  string   `json:"id"`
	Object              string   `json:"object"`
	ChargeEnabled       bool     `json:"charge_enabled"`
	Country             string   `json:"country"`
	CurrenciesSupported []string `json:"currencies_supported"`
	DefaultCurrency     string   `json:"default_currency"`
	DetailsSubmitted    bool     `json:"details_submitted"`
	TransferEnabled     bool     `json:"transfer_enabled"`
	DisplayName         string   `json:"display_name"`
	Email               string   `json:"email"`
	StatementDescriptor string   `json:"statement_descriptor"`
}

type AccountClient struct {
	client Client
}

// Retrieve loads a account.
//
// For more information: https://stripe.com/docs/api#retrieve_account
func (c *AccountClient) Retrieve() (*Account, error) {
	account := Account{}
	err := c.client.get("/account", nil, &account)
	return &account, err
}
