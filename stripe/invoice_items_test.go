package stripe

import (
  "strconv"
  "testing"
  "net/url"
  "github.com/bmizerany/assert"
)

func TestInvoiceItemCreate(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/invoiceitems", "invoice_items/invoice_item.json")
  params := InvoiceItemParams{}
  item, _ := client.InvoiceItems.Create(&params)
  assert.Equal(t, item.Id, "ii_123456789")
}

func TestInvoiceItemsRetrieve(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/invoiceitems/ii_123456789", "invoice_items/invoice_item.json")
  item, _ := client.InvoiceItems.Retrieve("ii_123456789")
  assert.Equal(t, item.Id, "ii_123456789")
}

func TestInvoiceItemsUpdate(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/invoiceitems/ii_123456789", "invoice_items/invoice_item.json")
  item, _ := client.InvoiceItems.Update("ii_123456789", new(InvoiceItemParams))
  assert.Equal(t, item.Id, "ii_123456789")
}

func TestInvoiceItemsDelete(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/invoiceitems/ii_123456789", "delete.json")
  res, _ := client.InvoiceItems.Delete("ii_123456789")
  assert.Equal(t, res.Deleted, true)
}

func TestInvoiceItemsAll(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/invoiceitems", "invoice_items/invoice_items.json")
  items, _ := client.InvoiceItems.All()
  assert.Equal(t, items.Count, 1)
  assert.Equal(t, items.Data[0].Id, "ii_123456789")
}


func TestInvoiceItemsAllWithFilters(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/invoiceitems", "invoice_items/invoice_items.json")
  items, _ := client.InvoiceItems.AllWithFilters(Filters{})
  assert.Equal(t, items.Count, 1)
  assert.Equal(t, items.Data[0].Id, "ii_123456789")
}

func TestParseInvoiceItemParams(t *testing.T) {
  params := InvoiceItemParams{
    Customer: "cus_123456789",
    Amount: 1000,
    Currency: "USD",
    Invoice: "in_123456789",
    Description: "Description",
    Metadata: Metadata{
      "foo": "bar",
    },
  }
  values := url.Values{}
  parseInvoiceItemParams(&params, &values)
  assert.Equal(t, values.Get("customer"), params.Customer)
  assert.Equal(t, values.Get("amount"), strconv.Itoa(params.Amount))
  assert.Equal(t, values.Get("currency"), params.Currency)
  assert.Equal(t, values.Get("invoice"), params.Invoice)
  assert.Equal(t, values.Get("description"), params.Description)
  assert.Equal(t, values.Get("metadata[foo]"), params.Metadata["foo"])
}
