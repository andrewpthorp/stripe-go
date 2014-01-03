package stripe

import (
	"net/url"
	"strconv"
)

type InvoiceItem struct {
	Id                 string   `json:"id"`
	Object             string   `json:"object"`
	Livemode           bool     `json:"livemode"`
	Amount             int64    `json:"amount"`
	Currency           string   `json:"currency"`
	Customer           string   `json:"customer"`
  Date               int64    `json:"date"`
  Proration          bool     `json:"proration"`
	Description        string   `json:"description"`
	Invoice            string   `json:"invoice"`
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
  err := delete("/invoiceitems/" + id, nil, &response)
  return &response, err
}

// List lists the first 10 invoice items. It calls ListCount with 10 as the
// count and 0 as the offset, which are the defaults in the Stripe API.
//
// For more information: https://stripe.com/docs/api#list_invoice_items
func (c *InvoiceItemClient) List() ([]*InvoiceItem, error) {
	return c.ListCount(10, 0)
}

// ListCount lists `count` invoice items starting at `offset`.
//
// For more information: https://stripe.com/docs/api#list_invoice_items
func (c *InvoiceItemClient) ListCount(count, offset int) ([]*InvoiceItem, error) {
	type items struct{ Data []*InvoiceItem }
	list := items{}

	params := url.Values{
		"count":  {strconv.Itoa(count)},
		"offset": {strconv.Itoa(offset)},
	}

	err := get("/invoiceitems", params, &list)
	return list.Data, err
}

// parseInvoiceItemParams takes a pointer to a InvoiceItemParams and a pointer
// to a url.Values. It iterates over everything in the InvoiceItemParams struct
// and Adds what is there to the url.Values.
func parseInvoiceItemParams(params *InvoiceItemParams, values *url.Values) {

  if params.Customer != "" {
    values.Add("customer", params.Customer)
  }

  if params.Amount != 0 {
    values.Add("amount", strconv.Itoa(params.Amount))
  }

  if params.Currency != "" {
    values.Add("currency", params.Currency)
  }

  if params.Invoice != "" {
    values.Add("invoice", params.Invoice)
  }

  if params.Description != "" {
    values.Add("description", params.Description)
  }

}
