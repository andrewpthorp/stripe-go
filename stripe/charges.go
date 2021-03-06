package stripe

import (
	"net/url"
)

type Charge struct {
	Id                 string   `json:"id"`
	Object             string   `json:"object"`
	Livemode           bool     `json:"livemode"`
	Amount             int64    `json:"amount"`
	Captured           bool     `json:"captured"`
	Card               *Card    `json:"card"`
	Created            int64    `json:"created"`
	Currency           string   `json:"currency"`
	Paid               bool     `json:"paid"`
	Refunded           bool     `json:"refunded"`
	Refunds            []Refund `json:"refunds"`
	AmountRefunded     int64    `json:"amount_refunded"`
	BalanceTransaction string   `json:"balance_transaction"`
	Customer           string   `json:"customer"`
	Description        string   `json:"description"`
	Dispute            *Dispute `json:"dispute"`
	FailureCode        string   `json:"failure_code"`
	FailureMessage     string   `json:"failure_message"`
	Invoice            string   `json:"invoice"`
	Metadata           Metadata `json:"metadata"`
}

type ChargeListResponse struct {
	ListResponse
	Data []Charge `json:"data"`
}

type ChargeClient struct {
	client Client
}

// Create creates a charge.
//
// For more information: https://stripe.com/docs/api#create_charge
func (c *ChargeClient) Create(params *ChargeParams) (*Charge, error) {
	charge := Charge{}
	values := url.Values{}
	parseChargeParams(params, &values)
	err := c.client.post("/charges", values, &charge)
	return &charge, err
}

// Capture captures an existing, uncaptured charge.
//
// For more information: https://stripe.com/docs/api#charge_capture
func (c *ChargeClient) Capture(id string, params *ChargeParams) (*Charge, error) {
	charge := Charge{}
	values := url.Values{}
	parseChargeParams(params, &values)
	err := c.client.post("/charges/"+id+"/capture", values, &charge)
	return &charge, err
}

// Retrieve loads a charge.
//
// For more information: https://stripe.com/docs/api#retrieve_charge
func (c *ChargeClient) Retrieve(id string) (*Charge, error) {
	charge := Charge{}
	err := c.client.get("/charges/"+id, nil, &charge)
	return &charge, err
}

// Update updates a charge.
//
// For more information: https://stripe.com/docs/api#update_charge
func (c *ChargeClient) Update(id string, params *ChargeParams) (*Charge, error) {
	charge := Charge{}
	values := url.Values{}
	parseChargeParams(params, &values)
	err := c.client.post("/charges/"+id, values, &charge)
	return &charge, err
}

// All lists the first 10 charges. It calls AllWithFilters with a blank Filters
// so all defaults are used.
//
// For more information: https://stripe.com/docs/api#list_charges
func (c *ChargeClient) All() (*ChargeListResponse, error) {
	return c.AllWithFilters(Filters{})
}

// AllWithFilters takes a Filters and applies all valid filters for the action.
//
// For more information: https://stripe.com/docs/api#list_charges
func (c *ChargeClient) AllWithFilters(filters Filters) (*ChargeListResponse, error) {
	response := ChargeListResponse{}
	values := url.Values{}
	addFiltersToValues([]string{"count", "offset", "customer"}, filters, &values)

	err := c.client.get("/charges", values, &response)
	return &response, err
}

// Refund refunds a charge.
//
// For more information: https://stripe.com/docs/api#refund_charge
func (c *ChargeClient) Refund(id string, params *RefundParams) (*Charge, error) {
	charge := Charge{}
	values := url.Values{}
	addParamsToValues(params, &values)
	err := c.client.post("/charges/"+id+"/refund", values, &charge)
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

	// Use parseMetaData from metadata.go to setup the metadata param
	if params.Metadata != nil {
		parseMetadata(params.Metadata, values)
	}

	addParamsToValues(params, values)
}
