package stripe

import (
	"github.com/bmizerany/assert"
	"net/url"
	"testing"
)

func TestParseBankAccountParams(t *testing.T) {
	params := BankAccountParams{
		Country:       "US",
		RoutingNumber: "111111111",
		AccountNumber: "1234567890",
	}
	values := url.Values{}
	parseBankAccountParams(&params, &values)
	assert.Equal(t, values.Get("bank_account[country]"), params.Country)
	assert.Equal(t, values.Get("bank_account[routing_number]"), params.RoutingNumber)
	assert.Equal(t, values.Get("bank_account[account_number]"), params.AccountNumber)
}
