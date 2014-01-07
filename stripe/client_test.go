package stripe

import (
	"fmt"
	"github.com/bmizerany/assert"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient("abc123")
	assert.Equal(t, _apiKey, "abc123")
	assert.Equal(t, _apiUrl, apiUrl)
	assert.Equal(t, reflect.TypeOf(client).Name(), "Client")
}

func TestNewClientWith(t *testing.T) {
	client := NewClientWith("http://foo.bar", "token")
	assert.Equal(t, _apiKey, "token")
	assert.Equal(t, _apiUrl, "http://foo.bar")
	assert.Equal(t, reflect.TypeOf(client).Name(), "Client")
}

func TestResourceClients(t *testing.T) {
	client := NewClient("abc123")
	assert.Equal(t, reflect.TypeOf(*client.Account).Name(), "AccountClient")
	assert.Equal(t, reflect.TypeOf(*client.ApplicationFees).Name(), "ApplicationFeeClient")
	assert.Equal(t, reflect.TypeOf(*client.Balance).Name(), "BalanceClient")
	assert.Equal(t, reflect.TypeOf(*client.Cards).Name(), "CardClient")
	assert.Equal(t, reflect.TypeOf(*client.Charges).Name(), "ChargeClient")
	assert.Equal(t, reflect.TypeOf(*client.Coupons).Name(), "CouponClient")
	assert.Equal(t, reflect.TypeOf(*client.Customers).Name(), "CustomerClient")
	assert.Equal(t, reflect.TypeOf(*client.Discounts).Name(), "DiscountClient")
	assert.Equal(t, reflect.TypeOf(*client.Disputes).Name(), "DisputeClient")
	assert.Equal(t, reflect.TypeOf(*client.Events).Name(), "EventClient")
	assert.Equal(t, reflect.TypeOf(*client.Invoices).Name(), "InvoiceClient")
	assert.Equal(t, reflect.TypeOf(*client.InvoiceItems).Name(), "InvoiceItemClient")
	assert.Equal(t, reflect.TypeOf(*client.Plans).Name(), "PlanClient")
	assert.Equal(t, reflect.TypeOf(*client.Recipients).Name(), "RecipientClient")
	assert.Equal(t, reflect.TypeOf(*client.Subscriptions).Name(), "SubscriptionClient")
	assert.Equal(t, reflect.TypeOf(*client.Tokens).Name(), "TokenClient")
	assert.Equal(t, reflect.TypeOf(*client.Transfers).Name(), "TransferClient")
}

func TestGet(t *testing.T) {
	setup()
	defer teardown()

	serveMux.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		fmt.Fprint(w, loadFixture("sample.json"))
	})

	var response struct{ Foo string }
	get("/get", nil, &response)
	assert.Equal(t, response.Foo, "bar")
}

func TestPost(t *testing.T) {
	setup()
	defer teardown()

	serveMux.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		fmt.Fprint(w, loadFixture("sample.json"))
	})

	var response struct{ Foo string }
	post("/post", nil, &response)
	assert.Equal(t, response.Foo, "bar")
}

func TestDelete(t *testing.T) {
	setup()
	defer teardown()

	serveMux.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "DELETE", r.Method)
		fmt.Fprint(w, loadFixture("sample.json"))
	})

	var response struct{ Foo string }
	delete("/delete", nil, &response)
	assert.Equal(t, response.Foo, "bar")
}

func TestRequest(t *testing.T) {
	setup()
	defer teardown()

	serveMux.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, r.Header.Get("Stripe-Version"), apiVersion)
		assert.Equal(t, r.Header.Get("User-Agent"), userAgent)
		fmt.Fprint(w, loadFixture("sample.json"))
	})

	serveMux.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		http.Error(w, loadFixture("errors/invalid_request_error.json"), http.StatusBadRequest)
	})

	// Success
	var response struct{ Foo string }
	request("GET", "/get", nil, &response)
	assert.Equal(t, response.Foo, "bar")

	// Error
	err := request("POST", "/error", nil, nil)
	assert.Equal(t, err.Error(), "An error occurred.")
}

func TestParseParamsGET(t *testing.T) {
	u, _ := url.Parse("http://www.stripe.com")
	params := url.Values{}
	params.Add("foo", "bar")
	reader := parseParams("GET", params, u)
	assert.Equal(t, u.RawQuery, "foo=bar")
	assert.Equal(t, reader, nil)
}

func TestParseParamsPOST(t *testing.T) {
	u, _ := url.Parse("http://www.stripe.com")
	params := url.Values{}
	params.Add("foo", "bar")
	reader := parseParams("POST", params, u)
	body, _ := ioutil.ReadAll(reader)
	assert.Equal(t, u.RawQuery, "")
	assert.Equal(t, string(body), "foo=bar")
}
