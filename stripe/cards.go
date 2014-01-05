package stripe

import (
	"net/url"
	"strconv"
)

type Card struct {
	Id                string `json:"id"`
	Object            string `json:"object"`
	ExpMonth          int64  `json:"exp_month"`
	ExpYear           int64  `json:"exp_year"`
	Fingerprint       string `json:"fingerprint"`
	Last4             string `json:"last4"`
	Type              string `json:"type"`
	AddressCity       string `json:"address_city"`
	AddressCountry    string `json:"address_country"`
	AddressLine1      string `json:"address_line1"`
	AddressLine1Check string `json:"address_line1_check"`
	AddressLine2      string `json:"address_line2"`
	AddressState      string `json:"address_state"`
	AddressZip        string `json:"address_zip"`
	AddressZipCheck   string `json:"address_zip_check"`
	Country           string `json:"country"`
	Customer          string `json:"customer"`
	CVCCheck          string `json:"cvc_check"`
	Name              string `json:"name"`
}

type CardListResponse struct {
	Object string  `json:"object"`
	Url    string  `json:"url"`
	Count  int     `json:"count"`
	Data   []*Card `json:"data"`
}

type CardClient struct{}

// Create creates a card for a customer.
//
// For more information: https://stripe.com/docs/api#create_card
func (c *CardClient) Create(customerId string, params *CardParams) (*Card, error) {
	card := Card{}
	values := url.Values{}
	parseCardParams(params, &values, true)
	err := post("/customers/"+customerId+"/cards", values, &card)
	return &card, err
}

// Retrieve loads a customers card.
//
// For more information: https://stripe.com/docs/api#retrieve_card
func (c *CardClient) Retrieve(customerId, id string) (*Card, error) {
	card := Card{}
	err := get("/customers/"+customerId+"/cards/"+id, nil, &card)
	return &card, err
}

// Update updates a customers card.
//
// For more information: https://stripe.com/docs/api#update_card
func (c *CardClient) Update(customerId, id string, params *CardParams) (*Card, error) {
	card := Card{}
	values := url.Values{}
	parseCardParams(params, &values, false)
	err := post("/customers/"+customerId+"/cards/"+id, values, &card)
	return &card, err
}

// Delete deletes a customers card.
//
// For more information: https://stripe.com/docs/api#delete_card
func (c *CardClient) Delete(customerId, id string) (*DeleteResponse, error) {
	response := DeleteResponse{}
	err := delete("/customers/"+customerId+"/cards/"+id, nil, &response)
	return &response, err
}

// All lists the first 10 cards for a customer. It calls AllWithFilters with
// a blank Filters so all defaults are used.
//
// For more information: https://stripe.com/docs/api#list_cards
func (c *CardClient) All(customerId string) (*CardListResponse, error) {
	return c.AllWithFilters(customerId, Filters{})
}

// AllWithFilters takes a Filters and applies all valid filters for the action.
//
// For more information: https://stripe.com/docs/api#list_cards
func (c *CardClient) AllWithFilters(customerId string, filters Filters) (*CardListResponse, error) {
	response := CardListResponse{}
	values := url.Values{}
	addFiltersToValues([]string{"count", "offset"}, filters, &values)
	err := get("/customers/"+customerId+"/cards", values, &response)
	return &response, err
}

// parseCardParams takes a pointer to a CardParams and a pointer to a
// url.Values. It iterates over everything in the CardParams struct and Adds
// what is there to the url.Values.
//
// If a Token is set on CardParams, that will be Added as "card" to the
// url.Values and the rest of the CardParams are ignored.
//
// The last argument, `includeRoot`, determines whether the values are added
// inside of a card[]. This is normally true for creates and false for updates.
func parseCardParams(params *CardParams, values *url.Values, includeRoot bool) {

	// If a token is passed, we are using that and not allowing a dictionary.
	if params.Token != "" {
		values.Add("card", params.Token)
		return
	}

	var prefix, suffix string

	if includeRoot {
		prefix = "card["
		suffix = "]"
	}

	if params.Number != "" {
		values.Add(prefix+"number"+suffix, params.Number)
	}

	if params.CVC != "" {
		values.Add(prefix+"cvc"+suffix, params.CVC)
	}

	if params.ExpMonth != 0 {
		values.Add(prefix+"exp_month"+suffix, strconv.Itoa(params.ExpMonth))
	}

	if params.ExpYear != 0 {
		values.Add(prefix+"exp_year"+suffix, strconv.Itoa(params.ExpYear))
	}

	if params.Name != "" {
		values.Add(prefix+"name"+suffix, params.Name)
	}

	if params.AddressLine1 != "" {
		values.Add(prefix+"address_line1"+suffix, params.AddressLine1)
	}

	if params.AddressLine2 != "" {
		values.Add(prefix+"address_line2"+suffix, params.AddressLine2)
	}

	if params.AddressCity != "" {
		values.Add(prefix+"address_city"+suffix, params.AddressCity)
	}

	if params.AddressZip != "" {
		values.Add(prefix+"address_zip"+suffix, params.AddressZip)
	}

	if params.AddressState != "" {
		values.Add(prefix+"address_state"+suffix, params.AddressState)
	}

	if params.AddressCountry != "" {
		values.Add(prefix+"address_country"+suffix, params.AddressCountry)
	}
}
