package stripe

import (
	"net/url"
)

type InvoiceItem struct {
	Id          string   `json:"id"`
	Object      string   `json:"object"`
	Livemode    bool     `json:"livemode"`
	Amount      int64    `json:"amount"`
	Currency    string   `json:"currency"`
	Customer    string   `json:"customer"`
	Date        int64    `json:"date"`
	Proration   bool     `json:"proration"`
	Description string   `json:"description"`
	Invoice     string   `json:"invoice"`
	Metadata    Metadata `json:"metadata"`
}

type InvoiceItemListResponse struct {
	Object string         `json:"object"`
	Url    string         `json:"url"`
	Count  int            `json:"count"`
	Data   []*InvoiceItem `json:"data"`
}

type InvoiceItemClient struct{}

// Create creates an invoice item.
//
// For more information: https://stripe.com/docs/api#create_invoice_item
func (c *InvoiceItemClient) Create(params *InvoiceItemParams) (*InvoiceItem, error) {
	item := InvoiceItem{}
	values := url.Values{}
	parseInvoiceItemParams(params, &values)
	err := post("/invoiceitems", values, &item)
	return &item, err
}

// Retrieve loads an invoice item.
//
// For more information: https://stripe.com/docs/api#retrieve_invoice_item
func (c *InvoiceItemClient) Retrieve(id string) (*InvoiceItem, error) {
	item := InvoiceItem{}
	err := get("/invoiceitems/"+id, nil, &item)
	return &item, err
}

// Update updates an invoice item.
//
// For more information: https://stripe.com/docs/api#update_invoice_item
func (c *InvoiceItemClient) Update(id string, params *InvoiceItemParams) (*InvoiceItem, error) {
	item := InvoiceItem{}
	values := url.Values{}
	parseInvoiceItemParams(params, &values)
	err := post("/invoiceitems/"+id, values, &item)
	return &item, err
}

// Delete deletes an invoice item.
//
// For more information: https://stripe.com/docs/api#delete_invoice_item
func (c *InvoiceItemClient) Delete(id string) (*DeleteResponse, error) {
	response := DeleteResponse{}
	err := delete("/invoiceitems/"+id, nil, &response)
	return &response, err
}

// All lists the first 10 invoice items. It calls AllWithFilters with a blank
// Filters so all defaults are used.
//
// For more information: https://stripe.com/docs/api#list_invoice_items
func (c *InvoiceItemClient) All() (*InvoiceItemListResponse, error) {
	return c.AllWithFilters(Filters{})
}

// AllWithFilters takes a Filters and applies all valid filters for the action.
//
// For more information: https://stripe.com/docs/api#list_invoice_items
func (c *InvoiceItemClient) AllWithFilters(filters Filters) (*InvoiceItemListResponse, error) {
	response := InvoiceItemListResponse{}
  values := url.Values{}
  addFiltersToValues([]string{"count", "offset"}, filters, &values)
	err := get("/invoiceitems", values, &response)
	return &response, err
}

// parseInvoiceItemParams takes a pointer to a InvoiceItemParams and a pointer
// to a url.Values. It iterates over everything in the InvoiceItemParams struct
// and Adds what is there to the url.Values.
func parseInvoiceItemParams(params *InvoiceItemParams, values *url.Values) {

	// Use parseMetaData from metadata.go to setup the metadata param
	if params.Metadata != nil {
		parseMetadata(params.Metadata, values)
	}

	addParamsToValues(params, values)
}
