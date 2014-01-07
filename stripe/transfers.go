package stripe

import (
	"net/url"
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
	Object string     `json:"object"`
	Url    string     `json:"url"`
	Count  int        `json:"count"`
	Data   []Transfer `json:"data"`
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
	err := post("/transfers/"+id+"/cancel", nil, &transfer)
	return &transfer, err
}

// All lists the first 10 transfers. It calls AllWithFilters with a blank
// Filters so all defaults are used.
//
// For more information: https://stripe.com/docs/api#list_transfers
func (c *TransferClient) All() (*TransferListResponse, error) {
	return c.AllWithFilters(Filters{})
}

// AllWithFilters takes a Filters and applies all valid filters for the action.
//
// For more information: https://stripe.com/docs/api#list_transfers
func (c *TransferClient) AllWithFilters(filters Filters) (*TransferListResponse, error) {
	response := TransferListResponse{}
	values := url.Values{}
	addFiltersToValues([]string{"count", "offset", "recipient", "status"}, filters, &values)
	err := get("/transfers", values, &response)
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
