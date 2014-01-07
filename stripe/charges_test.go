package stripe

import (
	"github.com/bmizerany/assert"
	"net/url"
	"strconv"
	"testing"
)

func TestChargeCreate(t *testing.T) {
	setup()
	defer teardown()
	handleWithJSON("/charges", "charges/charge.json")
	params := ChargeParams{}
	charge, _ := client.Charges.Create(&params)
	assert.Equal(t, charge.Id, "ch_123456789")
}

func TestChargesCapture(t *testing.T) {
	setup()
	defer teardown()
	handleWithJSON("/charges/ch_123456789/capture", "charges/charge.json")
	charge, _ := client.Charges.Capture("ch_123456789", new(ChargeParams))
	assert.Equal(t, charge.Id, "ch_123456789")
}

func TestChargesRetrieve(t *testing.T) {
	setup()
	defer teardown()
	handleWithJSON("/charges/ch_123456789", "charges/charge.json")
	charge, _ := client.Charges.Retrieve("ch_123456789")
	assert.Equal(t, charge.Id, "ch_123456789")
}

func TestChargesUpdate(t *testing.T) {
	setup()
	defer teardown()
	handleWithJSON("/charges/ch_123456789", "charges/charge.json")
	charge, _ := client.Charges.Update("ch_123456789", new(ChargeParams))
	assert.Equal(t, charge.Id, "ch_123456789")
}

func TestChargesAll(t *testing.T) {
	setup()
	defer teardown()
	handleWithJSON("/charges", "charges/charges.json")
	charges, _ := client.Charges.All()
	assert.Equal(t, charges.Count, 1)
	assert.Equal(t, charges.Data[0].Id, "ch_123456789")
}

func TestChargesAllWithFilters(t *testing.T) {
	setup()
	defer teardown()
	handleWithJSON("/charges", "charges/charges.json")
	charges, _ := client.Charges.AllWithFilters(Filters{})
	assert.Equal(t, charges.Count, 1)
	assert.Equal(t, charges.Data[0].Id, "ch_123456789")
}

func TestChargesRefund(t *testing.T) {
	setup()
	defer teardown()
	handleWithJSON("/charges/ch_123456789/refund", "charges/charge.json")
	charge, _ := client.Charges.Refund("ch_123456789", new(RefundParams))
	assert.Equal(t, charge.Id, "ch_123456789")
}

func TestParseChargeParams(t *testing.T) {
	params := ChargeParams{
		Amount:         2500,
		Currency:       "USD",
		Customer:       "cus_123456789",
		Description:    "Charge",
		DisableCapture: true,
		ApplicationFee: 100,
		CardParams: &CardParams{
			Number: "4242424242424242",
		},
		Metadata: Metadata{
			"foo": "bar",
		},
	}
	values := url.Values{}
	parseChargeParams(&params, &values)
	assert.Equal(t, values.Get("amount"), strconv.Itoa(params.Amount))
	assert.Equal(t, values.Get("currency"), params.Currency)
	assert.Equal(t, values.Get("customer"), params.Customer)
	assert.Equal(t, values.Get("description"), params.Description)
	assert.Equal(t, values.Get("capture"), "false")
	assert.Equal(t, values.Get("application_fee"), strconv.Itoa(params.ApplicationFee))
	assert.Equal(t, values.Get("card[number]"), params.CardParams.Number)
	assert.Equal(t, values.Get("metadata[foo]"), params.Metadata["foo"])
}
