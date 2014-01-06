package stripe

import (
  "strconv"
  "testing"
  "net/url"
  "github.com/bmizerany/assert"
)

func TestCustomerCreate(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/customers", "customers/customer.json")
  params := CustomerParams{}
  customer, _ := client.Customers.Create(&params)
  assert.Equal(t, customer.Id, "cus_123456789")
}

func TestCustomersRetrieve(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/customers/cus_123456789", "customers/customer.json")
  customer, _ := client.Customers.Retrieve("cus_123456789")
  assert.Equal(t, customer.Id, "cus_123456789")
}

func TestCustomersUpdate(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/customers/cus_123456789", "customers/customer.json")
  customer, _ := client.Customers.Update("cus_123456789", new(CustomerParams))
  assert.Equal(t, customer.Id, "cus_123456789")
}

func TestCustomersDelete(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/customers/cus_123456789", "delete.json")
  res, _ := client.Customers.Delete("cus_123456789")
  assert.Equal(t, res.Deleted, true)
}

func TestCustomersAll(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/customers", "customers/customers.json")
  customers, _ := client.Customers.All()
  assert.Equal(t, customers.Count, 1)
  assert.Equal(t, customers.Data[0].Id, "cus_123456789")
}


func TestCustomersAllWithFilters(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/customers", "customers/customers.json")
  customers, _ := client.Customers.AllWithFilters(Filters{})
  assert.Equal(t, customers.Count, 1)
  assert.Equal(t, customers.Data[0].Id, "cus_123456789")
}

func TestParseCustomerParams(t *testing.T) {
  params := CustomerParams{
    AccountBalance: 2000,
    Coupon: "coupon",
    Description: "description",
    Email: "apt@stripe.com",
    Plan: "plan",
    Quantity: 1,
    TrialEnd: 123456789,
    CardParams: &CardParams{
      Number: "4242424242424242",
    },
    Metadata: Metadata{
      "foo": "bar",
    },
  }
  values := url.Values{}
  parseCustomerParams(&params, &values)
  assert.Equal(t, values.Get("account_balance"), strconv.Itoa(params.AccountBalance))
  assert.Equal(t, values.Get("coupon"), params.Coupon)
  assert.Equal(t, values.Get("description"), params.Description)
  assert.Equal(t, values.Get("email"), params.Email)
  assert.Equal(t, values.Get("plan"), params.Plan)
  assert.Equal(t, values.Get("quantity"), strconv.Itoa(params.Quantity))
  assert.Equal(t, values.Get("trial_end"), strconv.Itoa(params.TrialEnd))
  assert.Equal(t, values.Get("card[number]"), params.CardParams.Number)
  assert.Equal(t, values.Get("metadata[foo]"), params.Metadata["foo"])
}
