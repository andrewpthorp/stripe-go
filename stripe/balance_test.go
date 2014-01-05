package stripe

import (
  "testing"
  "github.com/bmizerany/assert"
)

func TestBalanceRetrieve(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/balance", "balances/balance.json")
  balance, _ := client.Balance.Retrieve()
  assert.Equal(t, balance.Pending[0].Amount, int64(25000))
  assert.Equal(t, balance.Available[0].Amount, int64(0))
}

func TestBalanceRetrieveTransaction(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/balance/history/txn_123456789", "balances/balance_transaction.json")
  txn, _ := client.Balance.RetrieveTransaction("txn_123456789")
  assert.Equal(t, txn.Id, "txn_123456789")
}

func TestBalanceHistory(t *testing.T){
  setup()
  defer teardown()
  handleWithJSON("/balance/history", "balances/balance_history.json")
  history, _ := client.Balance.History()
  assert.Equal(t, history.Count, 1)
  assert.Equal(t, history.Data[0].Id, "txn_123456789")
}

func TestBalanceHistoryWithFilters(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/balance/history", "balances/balance_history.json")
  history, _ := client.Balance.HistoryWithFilters(Filters{})
  assert.Equal(t, history.Count, 1)
  assert.Equal(t, history.Data[0].Id, "txn_123456789")
}
