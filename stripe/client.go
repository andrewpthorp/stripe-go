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
	apiKey = ""
)

type Client struct {
	UserAgent string
	Cards     CardClient
	Charges   ChargeClient
	Coupons   CouponClient
	Customers CustomerClient
	Discounts DiscountClient
  Disputes  DisputeClient
	Plans     PlanClient
}

// NewClient returns a Client and allows us to access the resource clients.
func NewClient(key string) Client {
	apiKey = key

	return Client{
		UserAgent: userAgent,
		Cards:     CardClient{},
		Charges:   ChargeClient{},
		Coupons:   CouponClient{},
		Customers: CustomerClient{},
		Discounts: DiscountClient{},
    Disputes:  DisputeClient{},
		Plans:     PlanClient{},
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
	u, err := url.Parse(apiUrl + path)
	if err != nil {
		return err
	}

	// Much Authentication!
	u.User = url.User(apiKey)

	// Build and make HTTP Request.
	bodyReader := parseParams(method, params, u)
	req, err := http.NewRequest(method, u.String(), bodyReader)
	if err != nil {
		return err
	}
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
