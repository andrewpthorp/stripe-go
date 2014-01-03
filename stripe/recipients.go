package stripe

import (
	"net/url"
	"strconv"
)

type Recipient struct {
	Id            string  `json:"id"`
	Object        string  `json:"object"`
	Livemode      bool    `json:"livemode"`
	Created       int64   `json:"created"`
	Type          string  `json:"type"`
	//ActiveAccount Account `json:"account"`
	Description   string  `json:"description"`
	Email         string  `json:"email"`
	Name          string  `json:"name"`
}

// The RecipientClient is the receiver for most standard recipient related endpoints.
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

// List lists the first 10 recipients. It calls ListCount with 10 as the count
// and 0 as the offset, which are the defaults in the Stripe API.
//
// For more information: https://stripe.com/docs/api#list_recipients
func (c *RecipientClient) List() ([]*Recipient, error) {
	return c.ListCount(10, 0)
}

// ListCount lists `count` recipients starting at `offset`.
//
// For more information: https://stripe.com/docs/api#list_recipients
func (c *RecipientClient) ListCount(count, offset int) ([]*Recipient, error) {
	type recipients struct{ Data []*Recipient }
	list := recipients{}

	params := url.Values{
		"count":  {strconv.Itoa(count)},
		"offset": {strconv.Itoa(offset)},
	}

	err := get("/recipients", params, &list)
	return list.Data, err
}

// parseRecipientParams takes a pointer to a RecipientParams and a pointer to a url.Values,
// it iterates over everything in the RecipientParams struct and Adds what is there
// to the url.Values.
func parseRecipientParams(params *RecipientParams, values *url.Values) {

  if params.Name != "" {
    values.Add("name", params.Name)
  }

  if params.Type != "" {
    values.Add("type", params.Type)
  }

  if params.TaxId != "" {
    values.Add("tax_id", params.TaxId)
  }

  // TODO: Add BankAccountParams.

  if params.Email != "" {
    values.Add("email", params.Email)
  }

  if params.Description != "" {
    values.Add("description", params.Description)
  }

}
