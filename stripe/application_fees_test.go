package stripe

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestApplicationFeesRetrieve(t *testing.T) {
	setup()
	defer teardown()
	handleWithJSON("/application_fees/fee_123456789", "application_fees/application_fee.json")
	fee, _ := client.ApplicationFees.Retrieve("fee_123456789")
	assert.Equal(t, fee.Id, "fee_123456789")
}

func TestApplicationFeesRefund(t *testing.T) {
	setup()
	defer teardown()
	handleWithJSON("/application_fees/fee_123456789/refund", "application_fees/application_fee.json")
	fee, _ := client.ApplicationFees.Refund("fee_123456789", new(RefundParams))
	assert.Equal(t, fee.Id, "fee_123456789")
}

func TestApplicationFeesAll(t *testing.T) {
	setup()
	defer teardown()
	handleWithJSON("/application_fees", "application_fees/application_fees.json")
	fees, _ := client.ApplicationFees.All()
	assert.Equal(t, fees.Count, 1)
	assert.Equal(t, fees.Data[0].Id, "fee_123456789")
}

func TestApplicationFeesAllWithFilters(t *testing.T) {
	setup()
	defer teardown()
	handleWithJSON("/application_fees", "application_fees/application_fees.json")
	fees, _ := client.ApplicationFees.AllWithFilters(Filters{})
	assert.Equal(t, fees.Count, 1)
	assert.Equal(t, fees.Data[0].Id, "fee_123456789")
}
