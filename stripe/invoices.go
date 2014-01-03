package stripe

import (
	"net/url"
	"strconv"
)

type InvoiceLineItem struct {
	Id          string           `json:"id"`
	Object      string           `json:"object"`
	Livemode    bool             `json:"livemode"`
	Amount      int64            `json:"amount"`
	Currency    string           `json:"currency"`
	Period      map[string]int64 `json:"period"`
	Proration   bool             `json:"proration"`
	Type        string           `json:"type"`
	Description string           `json:"description"`
	Plan        *Plan            `json:"plan"`
	Quantity    int64            `json:"quantity"`
}

type Invoice struct {
	Id                 string    `json:"id"`
	Object             string    `json:"object"`
	Livemode           bool      `json:"livemode"`
	AmountDue          int64     `json:"amount_due"`
	AttemptCount       int64     `json:"attempt_count"`
	Attempted          bool      `json:"attempted"`
	Closed             bool      `json:"closed"`
	Currency           string    `json:"currency"`
	Customer           string    `json:"customer"`
	Date               int64     `json:"date"`
	Paid               bool      `json:"paid"`
	PeriodEnd          int64     `json:"period_end"`
	PeriodStart        int64     `json:"period_start"`
	StartingBalance    int64     `json:"starting_balance"`
	Subtotal           int64     `json:"subtotal"`
	Total              int64     `json:"total"`
	ApplicationFee     int64     `json:"application_fee"`
	Charge             string    `json:"charge"`
	Discount           *Discount `json:"discount"`
	EndingBalance      int64     `json:"ending_balance"`
	NextPaymentAttempt int64     `json:"next_payment_attempt"`
	//Lines              []InvoiceLineItem `json:"lines"`
}

// InvoiceLineItemListResponse is what is returned with a List request.
type InvoiceLineItemListResponse struct {
	Object string             `json:"object"`
	Url    string             `json:"url"`
	Count  int                `json:"count"`
	Data   []*InvoiceLineItem `json:"data"`
}

// InvoiceListResponse is what is returned with a List request.
type InvoiceListResponse struct {
	Object string     `json:"object"`
	Url    string     `json:"url"`
	Count  int        `json:"count"`
	Data   []*Invoice `json:"data"`
}

type InvoiceClient struct{}

// Create creates an invoice.
//
// For more information: https://stripe.com/docs/api#create_invoice
func (c *InvoiceClient) Create(params *InvoiceParams) (*Invoice, error) {
	invoice := Invoice{}
	values := url.Values{}
	parseInvoiceParams(params, &values)
	err := post("/invoices", values, &invoice)
	return &invoice, err
}

// Retrieve loads an invoice.
//
// For more information: https://stripe.com/docs/api#retrieve_invoice
func (c *InvoiceClient) Retrieve(id string) (*Invoice, error) {
	invoice := Invoice{}
	err := get("/invoices/"+id, nil, &invoice)
	return &invoice, err
}

// Update updates an invoice.
//
// For more information: https://stripe.com/docs/api#update_invoice
func (c *InvoiceClient) Update(id string, params *InvoiceParams) (*Invoice, error) {
	invoice := Invoice{}
	values := url.Values{}
	parseInvoiceParams(params, &values)
	err := post("/invoices/"+id, values, &invoice)
	return &invoice, err
}

// List lists the first 10 invoices. It calls ListCount with 10 as the count and
// 0 as the offset, which are the defaults in the Stripe API.
//
// For more information: https://stripe.com/docs/api#list_invoices
func (c *InvoiceClient) List() (*InvoiceListResponse, error) {
	return c.ListCount(10, 0)
}

// ListCount lists `count` invoices starting at `offset`.
//
// For more information: https://stripe.com/docs/api#list_cards
func (c *InvoiceClient) ListCount(count, offset int) (*InvoiceListResponse, error) {
	response := InvoiceListResponse{}

	params := url.Values{
		"count":  {strconv.Itoa(count)},
		"offset": {strconv.Itoa(offset)},
	}

	err := get("/invoices", params, &response)
	return &response, err
}

// Upcoming loads an upcoming invoice for a customer.
//
// For more information: https://stripe.com/docs/api#retrieve_customer_invoice
func (c *InvoiceClient) RetrieveUpcoming(customerId string) (*Invoice, error) {
	invoice := Invoice{}
	params := url.Values{
		"customer": {customerId},
	}
	err := get("/invoices/upcoming", params, &invoice)
	return &invoice, err
}

// Pay pays an invoice.
//
// For more information: https://stripe.com/docs/api#pay_invoice
func (c *InvoiceClient) Pay(id string) (*Invoice, error) {
	invoice := Invoice{}
	err := post("/invoices/"+id+"/pay", nil, &invoice)
	return &invoice, err
}

// RetrieveLines loads the first 10 line items for an invoice. It calls
// RetrieveLinesCount with 10 as the count and 0 as the offset, which are the
// defaults in the Stripe API.
//
// For more information: https://stripe.com/docs/api#invoice_lines
func (c *InvoiceClient) RetrieveLines(invoiceId string) (*InvoiceLineItemListResponse, error) {
	return c.RetrieveLinesCount(invoiceId, 10, 0)
}

// RetrieveLinesCount loads `count` invoice line items starting at `offset`.
//
// For more information: https://stripe.com/docs/api#invoice_lines
func (c *InvoiceClient) RetrieveLinesCount(invoiceId string, count, offset int) (*InvoiceLineItemListResponse, error) {
	response := InvoiceLineItemListResponse{}

	params := url.Values{
		"count":  {strconv.Itoa(count)},
		"offset": {strconv.Itoa(offset)},
	}

	err := get("/invoices/"+invoiceId+"/lines", params, &response)
	return &response, err
}

// parseInvoiceParams takes a pointer to an InvoiceParams and a pointer to a
// url.Values. It iterates over everything in the InvoiceParams struct and Adds
// what is there to the url.Values.
func parseInvoiceParams(params *InvoiceParams, values *url.Values) {

	if params.Customer != "" {
		values.Add("customer", params.Customer)
	}

	if params.ApplicationFee != 0 {
		values.Add("application_fee", strconv.Itoa(params.ApplicationFee))
	}

	if params.Closed {
		values.Add("closed", strconv.FormatBool(params.Closed))
	}

}
