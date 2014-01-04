package stripe

import (
	"net/url"
	"strconv"
)

type Transfer struct {
	Id                  string       `json:"id"`
	Object              string       `json:"object"`
	Livemode            bool         `json:"livemode"`
	Amount              int64        `json:"amount"`
	Currency            string       `json:"currency"`
	Date                int64        `json:"date"`
	Status              string       `json:"status"`
	Account             *BankAccount `json:"account"`
	BalanceTransaction  string       `json:"balance_transaction"`
	Description         string       `json:"description"`
	Recipient           string       `json:"recipient"`
	StatementDescriptor string       `json:"statement_descriptor"`
	Metadata            Metadata     `json:"metadata"`
}

type TransferListResponse struct {
	Object string      `json:"object"`
	Url    string      `json:"url"`
	Count  int         `json:"count"`
	Data   []*Transfer `json:"data"`
}

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
func (c *TransferClient) List() (*TransferListResponse, error) {
	return c.ListCount(10, 0)
}

// ListCount lists `count` transfers starting at `offset`.
//
// For more information: https://stripe.com/docs/api#list_transfers
func (c *TransferClient) ListCount(count, offset int) (*TransferListResponse, error) {
	response := TransferListResponse{}

	params := url.Values{
		"count":  {strconv.Itoa(count)},
		"offset": {strconv.Itoa(offset)},
	}

	err := get("/transfers", params, &response)
	return &response, err
}

// parseTransferParams takes a pointer to a TransferParams and a pointer to a url.Values,
// it iterates over everything in the TransferParams struct and Adds what is there
// to the url.Values.
func parseTransferParams(params *TransferParams, values *url.Values) {

	// Use parseMetaData from metadata.go to setup the metadata param
	if params.Metadata != nil {
		parseMetadata(params.Metadata, values)
	}

	addParamsToValues(params, values)
}
