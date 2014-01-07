package stripe

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	client          *http.Client
	apiKey          string
	apiUrl          string
	apiVersion      string
	userAgent       string
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
func NewClient(client *http.Client, apiKey string) Client {
	return NewClientWith(client, apiUrl, apiKey)
}

// NewClientWith returns a Client and allows us to access the resource clients.
func NewClientWith(client *http.Client, apiUrl, apiKey string) Client {
	c := Client{
		apiKey:     apiKey,
		apiUrl:     apiUrl,
		apiVersion: apiVersion,
		userAgent:  userAgent,
	}

	if client == nil {
		c.client = http.DefaultClient
	} else {
		c.client = client
	}

	c.Account = &AccountClient{client: c}
	c.ApplicationFees = &ApplicationFeeClient{client: c}
	c.Balance = &BalanceClient{client: c}
	c.Cards = &CardClient{client: c}
	c.Charges = &ChargeClient{client: c}
	c.Coupons = &CouponClient{client: c}
	c.Customers = &CustomerClient{client: c}
	c.Discounts = &DiscountClient{client: c}
	c.Disputes = &DisputeClient{client: c}
	c.Events = &EventClient{client: c}
	c.Invoices = &InvoiceClient{client: c}
	c.InvoiceItems = &InvoiceItemClient{client: c}
	c.Plans = &PlanClient{client: c}
	c.Recipients = &RecipientClient{client: c}
	c.Subscriptions = &SubscriptionClient{client: c}
	c.Tokens = &TokenClient{client: c}
	c.Transfers = &TransferClient{client: c}

	return c
}

// get is a shortcut to the underlying request, which sends an HTTP GET.
func (c *Client) get(path string, params url.Values, v interface{}) error {
	return c.request("GET", path, params, v)
}

// post is a shortcut to the underlying request, which sends an HTTP POST.
func (c *Client) post(path string, params url.Values, v interface{}) error {
	return c.request("POST", path, params, v)
}

func (c *Client) delete(path string, params url.Values, v interface{}) error {
	return c.request("DELETE", path, params, v)
}

// request is the method that actually delivers the HTTP Requests.
func (c *Client) request(method, path string, params url.Values, v interface{}) error {

	// Parse the URL, path, User, etc.
	u, err := url.Parse(c.apiUrl + path)
	if err != nil {
		return err
	}

	// Much Authentication!
	u.User = url.User(c.apiKey)

	// Build HTTP Request.
	bodyReader := parseParams(method, params, u)
	req, err := http.NewRequest(method, u.String(), bodyReader)
	if err != nil {
		return err
	}

	// Pin API Version, simplify maintenance.
	req.Header.Set("Stripe-Version", c.apiVersion)
	req.Header.Set("User-Agent", c.userAgent)

	// Send HTTP Request.
	res, err := c.client.Do(req)
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
