package stripe

import (
  "testing"
  "github.com/bmizerany/assert"
)

func TestAccountRetrieve(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/account", "accounts/account.json")
  account, _ := client.Account.Retrieve()

  assert.Equal(t, account.Id, "acct_abc123")
  assert.Equal(t, account.Object, "account")
  assert.Equal(t, account.ChargeEnabled, false)
  assert.Equal(t, account.Country, "US")
  assert.Equal(t, account.CurrenciesSupported, []string{"usd"})
  assert.Equal(t, account.DefaultCurrency, "usd")
  assert.Equal(t, account.DetailsSubmitted, false)
  assert.Equal(t, account.TransferEnabled, false)
  assert.Equal(t, account.DisplayName, "Stripe Account")
  assert.Equal(t, account.Email, "apt@stripe.com")
  assert.Equal(t, account.StatementDescriptor, "Statement Descriptor")
}
