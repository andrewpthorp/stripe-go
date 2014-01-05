package stripe

import "net/url"

type Token struct {
	Id          string       `json:"id"`
	Object      string       `json:"object"`
	Livemode    bool         `json:"livemode"`
	Created     int64        `json:"created"`
	Type        string       `json:"type"`
	Used        bool         `json:"used"`
	BankAccount *BankAccount `json:"bank_account"`
	Card        *Card        `json:"card"`
}

type TokenClient struct{}

// Create creates a token. This method will either create a Card Token or a
// Bank Account Token. If both CardParams and BankAccountParams are set, the
// BankAccountParams are ignored, giving preference to the CardParams.
//
// For more information:
// https://stripe.com/docs/api#create_card_token
// https://stripe.com/docs/api#create_bank_account_token
func (c *TokenClient) Create(params *TokenParams) (*Token, error) {
	token := Token{}
	values := url.Values{}
	parseTokenParams(params, &values)
	err := post("/tokens", values, &token)
	return &token, err
}

func (c *TokenClient) Retrieve(id string) (*Token, error) {
	token := Token{}
	err := get("/tokens/"+id, nil, &token)
	return &token, err
}

// parseTokenParams takes a pointer to a TokenParams and a pointer to a
// url.Values. It iterates over everything in the TokenParams struct and Adds
// what is there to the url.Values. If both CardParams and BankAccountParams
// are present, it ignores the BankAccountParams, giving preference to the
// CardParams.
func parseTokenParams(params *TokenParams, values *url.Values) {

	// Use parseCardParams/parseBankAccountParams (whichever is appropriate)
	if params.CardParams != nil {
		parseCardParams(params.CardParams, values, true)
	} else if params.BankAccountParams != nil {
		parseBankAccountParams(params.BankAccountParams, values)
	}

	addParamsToValues(params, values)
}
