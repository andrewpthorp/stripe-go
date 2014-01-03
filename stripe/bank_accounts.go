package stripe

import "net/url"

type BankAccount struct {
	Id          string `json:"id"`
	Object      string `json:"object"`
	BankName    string `json:"bank_name"`
	Last4       string `json:"last4"`
	Country     string `json:"country"`
	Currency    string `json:"currency"`
	Validated   bool   `json:"validated"`
	Verified    bool   `json:"verified"`
	Fingerprint string `json:"fingerprint"`
}

// parseBankAccountParams takes a pointer to a BankAccountParams and a pointer
// to a url.Values. It iterates over everything in the BankAccountParams struct
// and Adds what is there to the url.Values.
func parseBankAccountParams(params *BankAccountParams, values *url.Values) {

  if params.Country != "" {
    values.Add("bank_account[country]", params.Country)
  }

  if params.RoutingNumber != "" {
    values.Add("bank_account[routing_number]", params.RoutingNumber)
  }

  if params.AccountNumber != "" {
    values.Add("bank_account[account_number]", params.AccountNumber)
  }

}
