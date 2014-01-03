package stripe

import (
	"net/url"
	"strconv"
)

type Customer struct {
	Id             string        `json:"id"`
	Object         string        `json:"object"`
	Livemode       bool          `json:"livemode"`
	Created        int64         `json:"created"`
	AccountBalance int64         `json:"account_balance"`
	Currency       string        `json:"currency"`
	DefaultCard    string        `json:"default_card"`
	Delinquent     bool          `json:"delinquent"`
	Discount       *Discount     `json:"discount"`
	Email          string        `json:"email"`
	Subscription   *Subscription `json:"subscription"`
	Metadata       Metadata      `json:"metadata"`
}

// Delete deletes a customer.
//
// For more information: https://stripe.com/docs/api#delete_customer
func (c *Customer) Delete() (*DeleteResponse, error) {
	response := DeleteResponse{}
	err := delete("/customers/"+c.Id, nil, &response)
	return &response, err
}

// Update updates a customer.
//
// For more information: https://stripe.com/docs/api#update_customer
func (c *Customer) Update(params *CustomerParams) (*Customer, error) {
	values := url.Values{}
	parseCustomerParams(params, &values)
	err := post("/customers/"+c.Id, values, c)
	return c, err
}

// The CustomerClient is the receiver for most standard customer related endpoints.
type CustomerClient struct{}

// Create creates a customer.
//
// For more information: https://stripe.com/docs/api#create_customer
func (c *CustomerClient) Create(params *CustomerParams) (*Customer, error) {
	customer := Customer{}
	values := url.Values{}
	parseCustomerParams(params, &values)
	err := post("/customers", values, &customer)
	return &customer, err
}

// Retrieve loads a customer.
//
// For more information: https://stripe.com/docs/api#retrieve_customer
func (c *CustomerClient) Retrieve(id string) (*Customer, error) {
	customer := Customer{}
	err := get("/customers/"+id, nil, &customer)
	return &customer, err
}

// Update updates a customer.
//
// For more information: https://stripe.com/docs/api#update_customer
func (c *CustomerClient) Update(id string, params *CustomerParams) (*Customer, error) {
	customer := Customer{}
	values := url.Values{}
	parseCustomerParams(params, &values)
	err := post("/customers/"+id, values, &customer)
	return &customer, err
}

// Delete deletes a customer.
//
// For more information: https://stripe.com/docs/api#delete_customer
func (c *CustomerClient) Delete(id string) (*DeleteResponse, error) {
	response := DeleteResponse{}
	err := delete("/customers/"+id, nil, &response)
	return &response, err
}

// List lists the first 10 customers. It calls ListCount with 10 as the count
// and 0 as the offset, which are the defaults in the Stripe API.
//
// For more information: https://stripe.com/docs/api#list_customers
func (c *CustomerClient) List() ([]*Customer, error) {
	return c.ListCount(10, 0)
}

// ListCount lists `count` customers starting at `offset`.
//
// For more information: https://stripe.com/docs/api#list_customers
func (c *CustomerClient) ListCount(count, offset int) ([]*Customer, error) {
	type customers struct{ Data []*Customer }
	list := customers{}

	params := url.Values{
		"count":  {strconv.Itoa(count)},
		"offset": {strconv.Itoa(offset)},
	}

	err := get("/customers", params, &list)
	return list.Data, err
}

// parseCustomerParams takes a pointer to a CustomerParams and a pointer to a url.Values,
// it iterates over everything in the CustomerParams struct and Adds what is there
// to the url.Values.
func parseCustomerParams(params *CustomerParams, values *url.Values) {

	// Use parseCardParams from cards.go to setup the card param
	if params.CardParams != nil {
		parseCardParams(params.CardParams, values, true)
	}

	// Use parseMetaData from metadata.go to setup the metadata param
	if params.Metadata != nil {
		parseMetadata(params.Metadata, values)
	}

	if params.AccountBalance != 0 {
		values.Add("account_balance", strconv.Itoa(params.AccountBalance))
	}

	if params.Coupon != "" {
		values.Add("coupon", params.Coupon)
	}

	if params.Description != "" {
		values.Add("description", params.Description)
	}

	if params.Email != "" {
		values.Add("email", params.Email)
	}

	if params.Plan != "" {
		values.Add("plan", params.Plan)
	}

	if params.Quantity != 0 {
		values.Add("quantity", strconv.Itoa(params.Quantity))
	}

	if params.TrialEnd != 0 {
		values.Add("trial_end", strconv.Itoa(params.TrialEnd))
	}
}
