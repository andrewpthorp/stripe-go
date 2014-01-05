package stripe

import (
	"net/url"
)

type ApplicationFee struct {
	Id                 string    `json:"id"`
	Object             string    `json:"object"`
	Livemode           bool      `json:"livemode"`
	Account            string    `json:"account"`
	Amount             int64     `json:"amount"`
	Application        string    `json:"application"`
	BalanceTransaction string    `json:"balance_transaction"`
	Charge             string    `json:"charge"`
	Created            int64     `json:"created"`
	Currency           string    `json:"currency"`
	Refunded           bool      `json:"refunded"`
	Refunds            []*Refund `json:"refunds"`
	AmountRefunded     int64     `json:"amount_refunded"`
}

type ApplicationFeeListResponse struct {
	Object string            `json:"object"`
	Url    string            `json:"url"`
	Count  int               `json:"count"`
	Data   []*ApplicationFee `json:"data"`
}

type ApplicationFeeClient struct{}

// Retrieve loads an application fee.
//
// For more information: https://stripe.com/docs/api#retrieve_application_fee
func (c *ApplicationFeeClient) Retrieve(id string) (*ApplicationFee, error) {
	fee := ApplicationFee{}
	err := get("/application_fees/"+id, nil, &fee)
	return &fee, err
}

// Refund refunds an application fee.
//
// For more information: https://stripe.com/docs/api#refund_application_fee
func (c *ApplicationFeeClient) Refund(id string, params *RefundParams) (*ApplicationFee, error) {
	fee := ApplicationFee{}
	values := url.Values{}
	addParamsToValues(params, &values)
	err := post("/application_fees/"+id+"/refund", values, &fee)
	return &fee, err
}

// All lists the first 10 application fees. It calls AllWithFilters with a blank
// Filters so all defaults are used.
//
// For more information: https://stripe.com/docs/api#list_application_fees
func (c *ApplicationFeeClient) All() (*ApplicationFeeListResponse, error) {
	return c.AllWithFilters(Filters{})
}

// AllWithFilters takes a Filters and applies all valid filters for the action.
//
// For more information: https://stripe.com/docs/api#list_application_fees
func (c *ApplicationFeeClient) AllWithFilters(filters Filters) (*ApplicationFeeListResponse, error) {
	response := ApplicationFeeListResponse{}
	values := url.Values{}
	addFiltersToValues([]string{"count", "offset", "charge"}, filters, &values)
	err := get("/application_fees", values, &response)
	return &response, err
}
