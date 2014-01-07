package stripe

import (
	"net/url"
)

type Customer struct {
	Id             string            `json:"id"`
	Object         string            `json:"object"`
	Livemode       bool              `json:"livemode"`
	Created        int64             `json:"created"`
	AccountBalance int64             `json:"account_balance"`
	Cards          *CardListResponse `json:"cards"`
	Currency       string            `json:"currency"`
	DefaultCard    string            `json:"default_card"`
	Delinquent     bool              `json:"delinquent"`
	Discount       *Discount         `json:"discount"`
	Email          string            `json:"email"`
	Subscription   *Subscription     `json:"subscription"`
	Metadata       Metadata          `json:"metadata"`
}

type CustomerListResponse struct {
	ListResponse
	Data []Customer `json:"data"`
}

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

// All lists the first 10 customers. It calls AllWithFilters with a blank Filters
// so all defaults are used.
//
// For more information: https://stripe.com/docs/api#list_customers
func (c *CustomerClient) All() (*CustomerListResponse, error) {
	return c.AllWithFilters(Filters{})
}

// AllWithFilters takes a Filters and applies all valid filters for the action.
//
// For more information: https://stripe.com/docs/api#list_customers
func (c *CustomerClient) AllWithFilters(filters Filters) (*CustomerListResponse, error) {
	response := CustomerListResponse{}
	values := url.Values{}
	addFiltersToValues([]string{"count", "offset"}, filters, &values)
	err := get("/customers", values, &response)
	return &response, err
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

	addParamsToValues(params, values)
}
