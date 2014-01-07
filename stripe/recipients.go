package stripe

import (
	"net/url"
)

type Recipient struct {
	Id            string       `json:"id"`
	Object        string       `json:"object"`
	Livemode      bool         `json:"livemode"`
	Created       int64        `json:"created"`
	Type          string       `json:"type"`
	ActiveAccount *BankAccount `json:"active_account"`
	Description   string       `json:"description"`
	Email         string       `json:"email"`
	Name          string       `json:"name"`
	Metadata      Metadata     `json:"metadata"`
}

type RecipientListResponse struct {
	ListResponse
	Data []Recipient `json:"data"`
}

type RecipientClient struct{}

// Create creates a recipient.
//
// For more information: https://stripe.com/docs/api#create_recipient
func (c *RecipientClient) Create(params *RecipientParams) (*Recipient, error) {
	recipient := Recipient{}
	values := url.Values{}
	parseRecipientParams(params, &values)
	err := post("/recipients", values, &recipient)
	return &recipient, err
}

// Retrieve loads a recipient.
//
// For more information: https://stripe.com/docs/api#retrieve_recipient
func (c *RecipientClient) Retrieve(id string) (*Recipient, error) {
	recipient := Recipient{}
	err := get("/recipients/"+id, nil, &recipient)
	return &recipient, err
}

// Update updates a recipient.
//
// For more information: https://stripe.com/docs/api#update_recipient
func (c *RecipientClient) Update(id string, params *RecipientParams) (*Recipient, error) {
	recipient := Recipient{}
	values := url.Values{}
	parseRecipientParams(params, &values)
	err := post("/recipients/"+id, values, &recipient)
	return &recipient, err
}

// Delete deletes a recipient.
//
// For more information: https://stripe.com/docs/api/#delete_recipient
func (c *RecipientClient) Delete(id string) (*DeleteResponse, error) {
	response := DeleteResponse{}
	err := delete("/recipients/"+id, nil, &response)
	return &response, err
}

// All lists the first 10 recipients. It calls AllWithFilters with a blank
// Filters so all defaults are used.
//
// For more information: https://stripe.com/docs/api#list_recipients
func (c *RecipientClient) All() (*RecipientListResponse, error) {
	return c.AllWithFilters(Filters{})
}

// AllWithFilters takes a Filters and applies all valid filters for the action.
//
// For more information: https://stripe.com/docs/api#list_recipients
func (c *RecipientClient) AllWithFilters(filters Filters) (*RecipientListResponse, error) {
	response := RecipientListResponse{}
	values := url.Values{}
	addFiltersToValues([]string{"count", "offset", "verified"}, filters, &values)
	err := get("/recipients", values, &response)
	return &response, err
}

// parseRecipientParams takes a pointer to a RecipientParams and a pointer to a url.Values,
// it iterates over everything in the RecipientParams struct and Adds what is there
// to the url.Values.
func parseRecipientParams(params *RecipientParams, values *url.Values) {

	// Use parseBankAccountParams from bank_accounts.go to setup the bank_account
	// param
	if params.BankAccountParams != nil {
		parseBankAccountParams(params.BankAccountParams, values)
	}

	// Use parseMetaData from metadata.go to setup the metadata param
	if params.Metadata != nil {
		parseMetadata(params.Metadata, values)
	}

	addParamsToValues(params, values)
}
