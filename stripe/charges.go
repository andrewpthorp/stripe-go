package stripe

import (
	"net/url"
	"strconv"
)

type Charge struct {
	Id                 string   `json:"id"`
	Object             string   `json:"object"`
	Livemode           bool     `json:"livemode"`
	Amount             int64    `json:"amount"`
	Captured           bool     `json:"captured"`
	Card               Card     `json:"card"`
	Created            int64    `json:"created"`
	Currency           string   `json:"currency"`
	Paid               bool     `json:"paid"`
	Refunded           bool     `json:"refunded"`
	Refunds            []Refund `json:"refunds"`
	AmountRefunded     int64    `json:"amount_refunded"`
	BalanceTransaction string   `json:"balance_transaction"`
	Customer           string   `json:"customer"`
	Description        string   `json:"description"`
	Dispute            Dispute  `json:"dispute"`
	FailureCode        string   `json:"failure_code"`
	FailureMessage     string   `json:"failure_message"`
	Invoice            string   `json:"invoice"`
}

type ChargeClient struct{}

// Create creates a charge.
//
// For more information: https://stripe.com/docs/api#create_charge
func (c *ChargeClient) Create(params *ChargeParams) (*Charge, error) {
	charge := Charge{}
	values := url.Values{}
	parseChargeParams(params, &values)
	err := post("/charges", values, &charge)
	return &charge, err
}

// Retrieve loads a charge.
//
// For more information: https://stripe.com/docs/api#retrieve_charge
func (c *ChargeClient) Retrieve(id string) (*Charge, error) {
	charge := Charge{}
	err := get("/charges/"+id, nil, &charge)
	return &charge, err
}

// Update updates a charge.
//
// For more information: https://stripe.com/docs/api#update_charge
func (c *ChargeClient) Update(id string, params *ChargeParams) (*Charge, error) {
	charge := Charge{}
	values := url.Values{}
	parseChargeParams(params, &values)
	err := post("/charges/"+id, values, &charge)
	return &charge, err
}

// List lists the first 10 charges. It calls ListCount with 10 as the count and
// 0 as the offset, which are the defaults in the Stripe API.
//
// For more information: https://stripe.com/docs/api#list_charges
func (c *ChargeClient) List() ([]*Charge, error) {
	return c.ListCount(10, 0)
}

// ListCount lists `count` charges starting at `offset`.
//
// For more information: https://stripe.com/docs/api#list_charges
func (c *ChargeClient) ListCount(count, offset int) ([]*Charge, error) {
	type charges struct{ Data []*Charge }
	list := charges{}

	params := url.Values{
		"count":  {strconv.Itoa(count)},
		"offset": {strconv.Itoa(offset)},
	}

	err := get("/charges", params, &list)
	return list.Data, err
}

// Refund refunds a charge.
//
// For more information: https://stripe.com/docs/api#refund_charge
func (c *ChargeClient) Refund(id string, params *ChargeRefundParams) (*Charge, error) {
	charge := Charge{}
	values := url.Values{}

	if params.Amount != 0 {
		values.Add("amount", strconv.Itoa(params.Amount))
	}

	if params.RefundApplicationFee != false {
		values.Add("refund_application_fee", strconv.FormatBool(params.RefundApplicationFee))
	}

	err := post("/charges/"+id+"/refund", values, &charge)
	return &charge, err
}

// parseChargeParams takes a pointer to a ChargeParams and a pointer to a
// url.Values. It iterates over everything in the ChargeParams struct and Adds
// what is there to the url.Values.
func parseChargeParams(params *ChargeParams, values *url.Values) {

	// Use parseCardParams from cards.go to setup the card param
	if params.CardParams != nil {
		parseCardParams(params.CardParams, values, true)
	}

	if params.Amount != 0 {
		values.Add("amount", strconv.Itoa(params.Amount))
	}

	if params.Currency != "" {
		values.Add("currency", params.Currency)
	}

	if params.Customer != "" {
		values.Add("customer", params.Customer)
	}

	if params.Description != "" {
		values.Add("description", params.Description)
	}

	// TODO: What to do with Capture?

	if params.ApplicationFee != 0 {
		values.Add("application_fee", strconv.Itoa(params.ApplicationFee))
	}
}
