package stripe

type Refund struct {
	Object             string `json:"object"`
	Amount             int64  `json:"amount"`
	Created            int64  `json:"created"`
	Currency           string `json:"currency"`
	BalanceTransaction string `json:"balance_transaction"`
}
