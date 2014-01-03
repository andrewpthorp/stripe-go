package stripe

import (
	"net/url"
	"strconv"
)

type Fund struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

type Balance struct {
	Object    string  `json:"object"`
	Livemode  bool    `json:"livemode"`
	Available []*Fund `json:"available"`
	Pending   []*Fund `json:"pending"`
}

type FeeDetails struct {
	Amount      int64  `json:"amount"`
	Currency    string `json:"currency"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Application string `json:"application"`
}

type BalanceTransaction struct {
	Id          string        `json:"id"`
	Object      string        `json:"object"`
	Source      string        `json:"source"`
	Amount      int64         `json:"amount"`
	Currency    string        `json:"currency"`
	Net         int64         `json:"net"`
	Type        string        `json:"type"`
	Created     int64         `json:"created"`
	AvailableOn int64         `json:"available_on"`
	Status      string        `json:"status"`
	Fee         int64         `json:"fee"`
	FeeDetails  []*FeeDetails `json:"fee_details"`
}

type BalanceClient struct{}

// Retrieve loads a balance.
//
// For more information: https://stripe.com/docs/api#retrieve_balance
func (c *BalanceClient) Retrieve() (*Balance, error) {
	balance := Balance{}
	err := get("/balance", nil, &balance)
	return &balance, err
}

// RetrieveTransaction loads a balance transaction.
//
// For more information: https://stripe.com/docs/api#retrieve_balance_transaction
func (c *BalanceClient) RetrieveTransaction(id string) (*BalanceTransaction, error) {
	balanceTransaction := BalanceTransaction{}
	err := get("/balance/history/"+id, nil, &balanceTransaction)
	return &balanceTransaction, err
}

// History lists the first 10 balances in the balance history. It calls HistoryCount
// with 10 as the count and 0 as the offset, which are the defaults in the
// Stripe API.
//
// For more information: https://stripe.com/docs/api#balance_history
func (c *BalanceClient) History() ([]*BalanceTransaction, error) {
	return c.HistoryCount(10, 0)
}

// HistoryCount lists `count` balances in the balance history starting at `offset`.
//
// For more information: https://stripe.com/docs/api#balance_history
func (c *BalanceClient) HistoryCount(count, offset int) ([]*BalanceTransaction, error) {
	type balances struct{ Data []*BalanceTransaction }
	list := balances{}

	params := url.Values{
		"count":  {strconv.Itoa(count)},
		"offset": {strconv.Itoa(offset)},
	}

	err := get("/balance/history", params, &list)
	return list.Data, err
}
