package stripe

import (
	"github.com/bmizerany/assert"
	"net/url"
	"testing"
)

func TestTokenCreate(t *testing.T) {
	setup()
	defer teardown()
	handleWithJSON("/tokens", "tokens/token.json")
	params := TokenParams{}
	token, _ := client.Tokens.Create(&params)
	assert.Equal(t, token.Id, "tok_123456789")
}

func TestTokensRetrieve(t *testing.T) {
	setup()
	defer teardown()
	handleWithJSON("/tokens/tok_123456789", "tokens/token.json")
	token, _ := client.Tokens.Retrieve("tok_123456789")
	assert.Equal(t, token.Id, "tok_123456789")
}

func TestParseTokenParams(t *testing.T) {
	params := TokenParams{
		Customer: "cus_123456789",
		CardParams: &CardParams{
			Number: "4242424242424242",
		},
		BankAccountParams: &BankAccountParams{
			AccountNumber: "123456789",
		},
	}
	values := url.Values{}

	// Card Token
	parseTokenParams(&params, &values)
	assert.Equal(t, values.Get("customer"), params.Customer)
	assert.Equal(t, values.Get("card[number]"), params.CardParams.Number)
	assert.NotEqual(t, values.Get("bank_account[account_number]"), params.BankAccountParams.AccountNumber)

	// Bank Account Token
	params.CardParams = nil
	values = url.Values{}
	parseTokenParams(&params, &values)
	assert.Equal(t, values.Get("bank_account[account_number]"), params.BankAccountParams.AccountNumber)
}
