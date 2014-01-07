package stripe

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestDisputesUpdate(t *testing.T) {
	setup()
	defer teardown()
	handleWithJSON("/charges/ch_123456789/dispute", "disputes/dispute.json")
	dispute, _ := client.Disputes.Update("ch_123456789", "evidence")
	assert.Equal(t, dispute.Charge, "ch_123456789")
	assert.Equal(t, dispute.Amount, int64(1000))
	assert.Equal(t, dispute.Created, int64(123456789))
	assert.Equal(t, dispute.Status, "needs_response")
	assert.Equal(t, dispute.Livemode, false)
	assert.Equal(t, dispute.Currency, "usd")
	assert.Equal(t, dispute.Object, "dispute")
	assert.Equal(t, dispute.Reason, "general")
	assert.Equal(t, dispute.BalanceTransaction, "txn_123456789")
	assert.Equal(t, dispute.EvidenceDueBy, int64(123456789))
	assert.Equal(t, dispute.Evidence, "evidence")
}

func TestDisputesClose(t *testing.T) {
	setup()
	defer teardown()
	handleWithJSON("/charges/ch_123456789/dispute/close", "disputes/dispute.json")
	dispute, _ := client.Disputes.Close("ch_123456789")
	assert.Equal(t, dispute.Charge, "ch_123456789")
	assert.Equal(t, dispute.Amount, int64(1000))
	assert.Equal(t, dispute.Created, int64(123456789))
	assert.Equal(t, dispute.Status, "needs_response")
	assert.Equal(t, dispute.Livemode, false)
	assert.Equal(t, dispute.Currency, "usd")
	assert.Equal(t, dispute.Object, "dispute")
	assert.Equal(t, dispute.Reason, "general")
	assert.Equal(t, dispute.BalanceTransaction, "txn_123456789")
	assert.Equal(t, dispute.EvidenceDueBy, int64(123456789))
	assert.Equal(t, dispute.Evidence, "evidence")
}
