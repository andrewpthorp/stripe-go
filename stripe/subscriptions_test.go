package stripe

import (
	"github.com/bmizerany/assert"
	"net/url"
	"strconv"
	"testing"
)

func TestSubscriptionCreate(t *testing.T) {
	setup()
	defer teardown()
	handleWithJSON("/customers/cus_123456789/subscriptions", "subscriptions/subscription.json")
	params := SubscriptionParams{}
	subscription, _ := client.Subscriptions.Create("cus_123456789", &params)
  assert.Equal(t, subscription.Customer, "cus_123456789")
	assert.Equal(t, subscription.Id, "sub_123456789")
}

func TestSubscriptionsRetrieve(t *testing.T) {
	setup()
	defer teardown()
	handleWithJSON("/customers/cus_123456789/subscriptions/sub_123456789", "subscriptions/subscription.json")
	subscription, _ := client.Subscriptions.Retrieve("cus_123456789", "sub_123456789")
	assert.Equal(t, subscription.Id, "sub_123456789")
}

func TestSubscriptionsUpdate(t *testing.T) {
	setup()
	defer teardown()
	handleWithJSON("/customers/cus_123456789/subscriptions/sub_123456789", "subscriptions/subscription.json")
	subscription, _ := client.Subscriptions.Update("cus_123456789", "sub_123456789", new(SubscriptionParams))
	assert.Equal(t, subscription.Id, "sub_123456789")
}

func TestSubscriptionsDelete(t *testing.T) {
	setup()
	defer teardown()
	handleWithJSON("/customers/cus_123456789/subscriptions/sub_123456789", "subscriptions/subscription.json")
  params := SubscriptionParams{}
	subscription, _ := client.Subscriptions.Delete("cus_123456789", "sub_123456789", &params)
	assert.Equal(t, subscription.Id, "sub_123456789")
}

func TestSubscriptionsAll(t *testing.T) {
	setup()
	defer teardown()
	handleWithJSON("/customers/cus_123456789/subscriptions", "subscriptions/subscriptions.json")
	subscriptions, _ := client.Subscriptions.All("cus_123456789")
	assert.Equal(t, subscriptions.Count, 1)
	assert.Equal(t, subscriptions.Data[0].Id, "sub_123456789")
}

func TestSubscriptionsAllWithFilters(t *testing.T) {
	setup()
	defer teardown()
	handleWithJSON("/customers/cus_123456789/subscriptions", "subscriptions/subscriptions.json")
	subscriptions, _ := client.Subscriptions.AllWithFilters("cus_123456789", Filters{})
	assert.Equal(t, subscriptions.Count, 1)
	assert.Equal(t, subscriptions.Data[0].Id, "sub_123456789")
}

func TestParseSubscriptionParams(t *testing.T) {
	params := SubscriptionParams{
    Plan: "plan",
    Coupon: "coupon",
    DisableProrate: true,
    Quantity: 100,
    TrialEnd: 123456789,
    ApplicationFeePercent: 0.75,
		CardParams: &CardParams{
			Number: "4242424242424242",
		},
	}
	values := url.Values{}
	parseSubscriptionParams("create", &params, &values)
  assert.Equal(t, values.Get("plan"), params.Plan)
	assert.Equal(t, values.Get("coupon"), params.Coupon)
  assert.Equal(t, values.Get("prorate"), "false")
	assert.Equal(t, values.Get("quantity"), strconv.Itoa(params.Quantity))
	assert.Equal(t, values.Get("trial_end"), strconv.Itoa(params.TrialEnd))
  assert.Equal(t, values.Get("application_fee_percent"), "0.75")
	assert.Equal(t, values.Get("card[number]"), params.CardParams.Number)
}
