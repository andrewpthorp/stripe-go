package stripe

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var (
	_apiKey = ""
	_apiUrl = ""
)

type Client struct {
	Account         *AccountClient
	ApplicationFees *ApplicationFeeClient
	Balance         *BalanceClient
	Cards           *CardClient
	Charges         *ChargeClient
	Coupons         *CouponClient
	Customers       *CustomerClient
	Discounts       *DiscountClient
	Disputes        *DisputeClient
	Events          *EventClient
	Invoices        *InvoiceClient
	InvoiceItems    *InvoiceItemClient
	Plans           *PlanClient
	Recipients      *RecipientClient
	Subscriptions   *SubscriptionClient
	Tokens          *TokenClient
	Transfers       *TransferClient
}

// NewClient returns a Client and sets the apiUrl to the live apiUrl.
func NewClient(apiKey string) Client {
	return NewClientWith(apiUrl, apiKey)
}

// NewClientWith returns a Client and allows us to access the resource clients.
func NewClientWith(apiUrl, apiKey string) Client {
	_apiKey = apiKey
	_apiUrl = apiUrl

	return Client{
		Account:         new(AccountClient),
		ApplicationFees: new(ApplicationFeeClient),
		Balance:         new(BalanceClient),
		Cards:           new(CardClient),
		Charges:         new(ChargeClient),
		Coupons:         new(CouponClient),
		Customers:       new(CustomerClient),
		Discounts:       new(DiscountClient),
		Disputes:        new(DisputeClient),
		Events:          new(EventClient),
		Invoices:        new(InvoiceClient),
		InvoiceItems:    new(InvoiceItemClient),
		Plans:           new(PlanClient),
		Recipients:      new(RecipientClient),
		Subscriptions:   new(SubscriptionClient),
		Tokens:          new(TokenClient),
		Transfers:       new(TransferClient),
	}
}

// get is a shortcut to the underlying request, which sends an HTTP GET.
func get(path string, params url.Values, v interface{}) error {
	return request("GET", path, params, v)
}

// post is a shortcut to the underlying request, which sends an HTTP POST.
func post(path string, params url.Values, v interface{}) error {
	return request("POST", path, params, v)
}

func delete(path string, params url.Values, v interface{}) error {
	return request("DELETE", path, params, v)
}

// request is the method that actually delivers the HTTP Requests.
func request(method, path string, params url.Values, v interface{}) error {

	// Parse the URL, path, User, etc.
	u, err := url.Parse(_apiUrl + path)
	if err != nil {
		return err
	}

	// Much Authentication!
	u.User = url.User(_apiKey)

	// Build HTTP Request.
	bodyReader := parseParams(method, params, u)
	req, err := http.NewRequest(method, u.String(), bodyReader)
	if err != nil {
		return err
	}

	// Pin API Version, simplify maintenance.
	req.Header.Set("Stripe-Version", apiVersion)
	req.Header.Set("User-Agent", userAgent)

	// Send HTTP Request.
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	// Read response.
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return err
	}

	// If the API didn't return a 200, parse the error and return it.
	if res.StatusCode != 200 {
		err := ErrorResponse{}
		json.Unmarshal(body, &err)
		return &err
	}

	// Parse the body, store it in v, return the result of Unmarshal.
	return json.Unmarshal(body, v)
}

// parseParams takes a method, url.Values and a pointer to a url.URL. If the
// method is "GET", it adds the encoded url.Values to the rawQuery of the
// url.URL. If the method is not "GET", it creates a new io.Reader from the
// encoded url.Values and returns them.
func parseParams(method string, params url.Values, url *url.URL) io.Reader {
	var reader io.Reader
	encoded := params.Encode()

	switch method {
	case "GET":
		url.RawQuery = encoded
	default:
		reader = strings.NewReader(encoded)
	}

	return reader
}
