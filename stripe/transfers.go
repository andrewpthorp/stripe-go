package stripe

import (
	"net/url"
	"strconv"
)

type Transfer struct {
	Id                  string  `json:"id"`
	Object              string  `json:"object"`
	Livemode            bool    `json:"livemode"`
	Amount              int64   `json:"amount"`
	Currency            string  `json:"currency"`
	Date                int64   `json:"date"`
	Status              string  `json:"status"`
	//Account             Account `json:"account"`
	BalanceTransaction  string  `json:"balance_transaction"`
	Description         string  `json:"description"`
	Recipient           string  `json:"recipient"`
	StatementDescriptor string  `json:"statement_descriptor"`
}

// The TransferClient is the receiver for most standard transfer related endpoints.
type TransferClient struct{}

// Create creates a transfer.
//
// For more information: https://stripe.com/docs/api#create_transfer
func (c *TransferClient) Create(params *TransferParams) (*Transfer, error) {
	transfer := Transfer{}
	values := url.Values{}
	parseTransferParams(params, &values)
	err := post("/transfers", values, &transfer)
	return &transfer, err
}

// Retrieve loads a transfer.
//
// For more information: https://stripe.com/docs/api#retrieve_transfer
func (c *TransferClient) Retrieve(id string) (*Transfer, error) {
	transfer := Transfer{}
	err := get("/transfers/"+id, nil, &transfer)
	return &transfer, err
}

// Update updates a transfer.
//
// For more information: https://stripe.com/docs/api#update_transfer
func (c *TransferClient) Update(id string, params *TransferParams) (*Transfer, error) {
	transfer := Transfer{}
	values := url.Values{}
	parseTransferParams(params, &values)
	err := post("/transfers/"+id, values, &transfer)
	return &transfer, err
}

// Cancel cancels a transfer.
//
// For more information: https://stripe.com/docs/api/#cancel_transfer
func (c *TransferClient) Cancel(id string) (*Transfer, error) {
  transfer := Transfer{}
  err := post("/transfers/"+id, nil, &transfer)
  return &transfer, err
}

// List lists the first 10 transfers. It calls ListCount with 10 as the count
// and 0 as the offset, which are the defaults in the Stripe API.
//
// For more information: https://stripe.com/docs/api#list_transfers
func (c *TransferClient) List() ([]*Transfer, error) {
	return c.ListCount(10, 0)
}

// ListCount lists `count` transfers starting at `offset`.
//
// For more information: https://stripe.com/docs/api#list_transfers
func (c *TransferClient) ListCount(count, offset int) ([]*Transfer, error) {
	type transfers struct{ Data []*Transfer }
	list := transfers{}

	params := url.Values{
		"count":  {strconv.Itoa(count)},
		"offset": {strconv.Itoa(offset)},
	}

	err := get("/transfers", params, &list)
	return list.Data, err
}

// parseTransferParams takes a pointer to a TransferParams and a pointer to a url.Values,
// it iterates over everything in the TransferParams struct and Adds what is there
// to the url.Values.
func parseTransferParams(params *TransferParams, values *url.Values) {

  if params.Amount != 0 {
    values.Add("amount", strconv.Itoa(params.Amount))
  }

  if params.Currency != "" {
    values.Add("currency", params.Currency)
  }

  if params.Recipient != "" {
    values.Add("recipient", params.Recipient)
  }

  if params.Description != "" {
    values.Add("description", params.Description)
  }

  if params.StatementDescriptor != "" {
    values.Add("statement_descriptor", params.StatementDescriptor)
  }

}
