package stripe

import (
  "strconv"
  "testing"
  "net/url"
  "github.com/bmizerany/assert"
)

func TestTransferCreate(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/transfers", "transfers/transfer.json")
  params := TransferParams{}
  transfer, _ := client.Transfers.Create(&params)
  assert.Equal(t, transfer.Id, "tr_123456789")
}

func TestTransfersRetrieve(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/transfers/tr_123456789", "transfers/transfer.json")
  transfer, _ := client.Transfers.Retrieve("tr_123456789")
  assert.Equal(t, transfer.Id, "tr_123456789")
}

func TestTransfersUpdate(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/transfers/tr_123456789", "transfers/transfer.json")
  transfer, _ := client.Transfers.Update("tr_123456789", new(TransferParams))
  assert.Equal(t, transfer.Id, "tr_123456789")
}

func TestTransfersCancel(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/transfers/tr_123456789/cancel", "transfers/transfer.json")
  transfer, _ := client.Transfers.Cancel("tr_123456789")
  assert.Equal(t, transfer.Id, "tr_123456789")
}

func TestTransfersAll(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/transfers", "transfers/transfers.json")
  transfers, _ := client.Transfers.All()
  assert.Equal(t, transfers.Count, 1)
  assert.Equal(t, transfers.Data[0].Id, "tr_123456789")
}


func TestTransfersAllWithFilters(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/transfers", "transfers/transfers.json")
  transfers, _ := client.Transfers.AllWithFilters(Filters{})
  assert.Equal(t, transfers.Count, 1)
  assert.Equal(t, transfers.Data[0].Id, "tr_123456789")
}

func TestParseTransferParams(t *testing.T) {
  params := TransferParams{
    Amount: 1000,
    Currency: "USD",
    Recipient: "rp_123456789",
    Description: "Description",
    StatementDescriptor: "Statement Descriptor",
    Metadata: Metadata{
      "foo": "bar",
    },
  }
  values := url.Values{}
  parseTransferParams(&params, &values)
  assert.Equal(t, values.Get("amount"), strconv.Itoa(params.Amount))
  assert.Equal(t, values.Get("currency"), params.Currency)
  assert.Equal(t, values.Get("recipient"), params.Recipient)
  assert.Equal(t, values.Get("description"), params.Description)
  assert.Equal(t, values.Get("statement_descriptor"), params.StatementDescriptor)
  assert.Equal(t, values.Get("metadata[foo]"), params.Metadata["foo"])
}
