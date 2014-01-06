package stripe

import (
  "strconv"
  "testing"
  "net/url"
  "github.com/bmizerany/assert"
)

func TestInvoiceCreate(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/invoices", "invoices/invoice.json")
  params := InvoiceParams{}
  invoice, _ := client.Invoices.Create(&params)
  assert.Equal(t, invoice.Id, "in_123456789")
}

func TestInvoicesRetrieve(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/invoices/in_123456789", "invoices/invoice.json")
  invoice, _ := client.Invoices.Retrieve("in_123456789")
  assert.Equal(t, invoice.Id, "in_123456789")
}

func TestInvoicesUpdate(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/invoices/in_123456789", "invoices/invoice.json")
  invoice, _ := client.Invoices.Update("in_123456789", new(InvoiceParams))
  assert.Equal(t, invoice.Id, "in_123456789")
}

func TestInvoicesAll(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/invoices", "invoices/invoices.json")
  invoices, _ := client.Invoices.All()
  assert.Equal(t, invoices.Count, 1)
  assert.Equal(t, invoices.Data[0].Id, "in_123456789")
}

func TestInvoicesAllWithFilters(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/invoices", "invoices/invoices.json")
  invoices, _ := client.Invoices.AllWithFilters(Filters{})
  assert.Equal(t, invoices.Count, 1)
  assert.Equal(t, invoices.Data[0].Id, "in_123456789")
}

func TestInvoicesRetrieveUpcoming(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/invoices/upcoming", "invoices/invoice.json")
  invoice, _ := client.Invoices.RetrieveUpcoming("cus_123456789")
  assert.Equal(t, invoice.Id, "in_123456789")
}

func TestInvoicesPay(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/invoices/in_123456789/pay", "invoices/invoice.json")
  invoice, _ := client.Invoices.Pay("in_123456789")
  assert.Equal(t, invoice.Id, "in_123456789")
}

func TestInvoicesRetrieveLines(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/invoices/in_123456789/lines", "invoices/lines.json")
  invoices, _ := client.Invoices.RetrieveLines("in_123456789")
  assert.Equal(t, invoices.Count, 1)
  assert.Equal(t, invoices.Data[0].Id, "ii_123456789")
}

func TestInvoicesRetrieveLinesWithFilters(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/invoices/in_123456789/lines", "invoices/lines.json")
  invoices, _ := client.Invoices.RetrieveLinesWithFilters("in_123456789", Filters{})
  assert.Equal(t, invoices.Count, 1)
  assert.Equal(t, invoices.Data[0].Id, "ii_123456789")
}

func TestParseInvoiceParams(t *testing.T) {
  params := InvoiceParams{
    Customer: "cus_123456789",
    ApplicationFee: 2500,
    Closed: true,
  }
  values := url.Values{}
  parseInvoiceParams(&params, &values)
  assert.Equal(t, values.Get("customer"), params.Customer)
  assert.Equal(t, values.Get("application_fee"), strconv.Itoa(params.ApplicationFee))
  assert.Equal(t, values.Get("closed"), strconv.FormatBool(params.Closed))
}
