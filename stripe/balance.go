package stripe

import (
	"net/url"
)

type Fund struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

type Balance struct {
	Object    string `json:"object"`
	Livemode  bool   `json:"livemode"`
	Available []Fund `json:"available"`
	Pending   []Fund `json:"pending"`
}

type FeeDetails struct {
	Amount      int64  `json:"amount"`
	Currency    string `json:"currency"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Application string `json:"application"`
}

type BalanceTransaction struct {
	Id          string       `json:"id"`
	Object      string       `json:"object"`
	Source      string       `json:"source"`
	Amount      int64        `json:"amount"`
	Currency    string       `json:"currency"`
	Net         int64        `json:"net"`
	Type        string       `json:"type"`
	Created     int64        `json:"created"`
	AvailableOn int64        `json:"available_on"`
	Status      string       `json:"status"`
	Fee         int64        `json:"fee"`
	FeeDetails  []FeeDetails `json:"fee_details"`
}

type BalanceTransactionListResponse struct {
	ListResponse
	Data []BalanceTransaction `json:"data"`
}

type BalanceClient struct {
	client Client
}

// Retrieve loads a balance.
//
// For more information: https://stripe.com/docs/api#retrieve_balance
func (c *BalanceClient) Retrieve() (*Balance, error) {
	balance := Balance{}
	err := c.client.get("/balance", nil, &balance)
	return &balance, err
}

// RetrieveTransaction loads a balance transaction.
//
// For more information: https://stripe.com/docs/api#retrieve_balance_transaction
func (c *BalanceClient) RetrieveTransaction(id string) (*BalanceTransaction, error) {
	balanceTransaction := BalanceTransaction{}
	err := c.client.get("/balance/history/"+id, nil, &balanceTransaction)
	return &balanceTransaction, err
}

// History lists the first 10 balances in the balance history. It calls
// HistoryWithFilters with a blank Filters so all defaults are used.
//
// For more information: https://stripe.com/docs/api#balance_history
func (c *BalanceClient) History() (*BalanceTransactionListResponse, error) {
	return c.HistoryWithFilters(Filters{})
}

// HistoryWithFilters takes a Filters and applies all valid filters for the action.
//
// For more information: https://stripe.com/docs/api#balance_history
func (c *BalanceClient) HistoryWithFilters(filters Filters) (*BalanceTransactionListResponse, error) {
	response := BalanceTransactionListResponse{}
	values := url.Values{}
	addFiltersToValues([]string{"count", "offset", "currency", "source", "transfer", "type"}, filters, &values)
	err := c.client.get("/balance/history", values, &response)
	return &response, err
}
