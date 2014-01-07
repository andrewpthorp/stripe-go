package stripe

import (
	"github.com/bmizerany/assert"
	"net/url"
	"testing"
)

func TestRecipientCreate(t *testing.T) {
	setup()
	defer teardown()
	handleWithJSON("/recipients", "recipients/recipient.json")
	params := RecipientParams{}
	recipient, _ := client.Recipients.Create(&params)
	assert.Equal(t, recipient.Id, "rp_123456789")
}

func TestRecipientsRetrieve(t *testing.T) {
	setup()
	defer teardown()
	handleWithJSON("/recipients/rp_123456789", "recipients/recipient.json")
	recipient, _ := client.Recipients.Retrieve("rp_123456789")
	assert.Equal(t, recipient.Id, "rp_123456789")
}

func TestRecipientsUpdate(t *testing.T) {
	setup()
	defer teardown()
	handleWithJSON("/recipients/rp_123456789", "recipients/recipient.json")
	recipient, _ := client.Recipients.Update("rp_123456789", new(RecipientParams))
	assert.Equal(t, recipient.Id, "rp_123456789")
}

func TestRecipientsDelete(t *testing.T) {
	setup()
	defer teardown()
	handleWithJSON("/recipients/rp_123456789", "delete.json")
	res, _ := client.Recipients.Delete("rp_123456789")
	assert.Equal(t, res.Deleted, true)
}

func TestRecipientsAll(t *testing.T) {
	setup()
	defer teardown()
	handleWithJSON("/recipients", "recipients/recipients.json")
	recipients, _ := client.Recipients.All()
	assert.Equal(t, recipients.Count, 1)
	assert.Equal(t, recipients.Data[0].Id, "rp_123456789")
}

func TestRecipientsAllWithFilters(t *testing.T) {
	setup()
	defer teardown()
	handleWithJSON("/recipients", "recipients/recipients.json")
	recipients, _ := client.Recipients.AllWithFilters(Filters{})
	assert.Equal(t, recipients.Count, 1)
	assert.Equal(t, recipients.Data[0].Id, "rp_123456789")
}

func TestParseRecipientParams(t *testing.T) {
	params := RecipientParams{
		Name:        "Andrew Thorp",
		Type:        "individual",
		TaxId:       "123456789",
		Email:       "apt@stripe.com",
		Description: "Description",
		BankAccountParams: &BankAccountParams{
			AccountNumber: "123456789",
		},
		Metadata: Metadata{
			"foo": "bar",
		},
	}
	values := url.Values{}
	parseRecipientParams(&params, &values)
	assert.Equal(t, values.Get("name"), params.Name)
	assert.Equal(t, values.Get("type"), params.Type)
	assert.Equal(t, values.Get("tax_id"), params.TaxId)
	assert.Equal(t, values.Get("email"), params.Email)
	assert.Equal(t, values.Get("description"), params.Description)
	assert.Equal(t, values.Get("bank_account[account_number]"), params.BankAccountParams.AccountNumber)
	assert.Equal(t, values.Get("metadata[foo]"), params.Metadata["foo"])
}
