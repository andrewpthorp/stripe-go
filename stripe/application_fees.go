package stripe

import (
	"net/url"
	"strconv"
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

// ApplicationFeeListResponse is what is returned with a List request.
type ApplicationFeeListResponse struct {
	Object string            `json:"object"`
	Url    string            `json:"url"`
	Count  int               `json:"count"`
	Data   []*ApplicationFee `json:"data"`
}

// ApplicationFeeClient is the receiver for all ApplicationFee methods of the
// API.
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

	if params.Amount != 0 {
		values.Add("amount", strconv.Itoa(params.Amount))
	}

	err := post("/application_fees/"+id+"/refund", values, &fee)
	return &fee, err
}

// List lists the first 10 application fees. It calls ListCount with 10 as the
// count and 0 as the offset, which are the defaults in the Stripe API.
//
// For more information: https://stripe.com/docs/api#list_application_fees
func (c *ApplicationFeeClient) List() (*ApplicationFeeListResponse, error) {
	return c.ListCount(10, 0)
}

// ListCount lists `count` application fees starting at `offset`.
//
// For more information: https://stripe.com/docs/api#list_application_fees
func (c *ApplicationFeeClient) ListCount(count, offset int) (*ApplicationFeeListResponse, error) {
	response := ApplicationFeeListResponse{}

	params := url.Values{
		"count":  {strconv.Itoa(count)},
		"offset": {strconv.Itoa(offset)},
	}

	err := get("/application_fees", params, &response)
	return &response, err
}
