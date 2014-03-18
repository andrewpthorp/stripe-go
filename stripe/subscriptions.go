package stripe

import (
	"net/url"
	"strconv"
)

type Subscription struct {
	Id                    string  `json:"id"`
	Object                string  `json:"object"`
	CancelAtPeriodEnd     bool    `json:"cancel_at_period_end"`
	Customer              string  `json:"customer"`
	Plan                  *Plan   `json:"plan"`
	Quantity              int64   `json:"quantity"`
	Start                 int64   `json:"start"`
	Status                string  `json:"status"`
	ApplicationFeePercent float64 `json:"application_fee_percent"`
	CanceledAt            int64   `json:"canceled_at"`
	CurrentPeriodEnd      int64   `json:"current_period_end"`
	CurrentPeriodStart    int64   `json:"current_period_start"`
	EndedAt               int64   `json:"ended_at"`
	TrialEnd              int64   `json:"trial_end"`
	TrialStart            int64   `json:"trial_start"`
}

type SubscriptionListResponse struct {
  ListResponse
  Data []Subscription `json:"data"`
}

type SubscriptionClient struct {
	client Client
}

// Create creates a subscription.
//
// For more information: https://stripe.com/docs/api#create_subscription
func (c *SubscriptionClient) Create(customerId string, params *SubscriptionParams) (*Subscription, error) {
	subscription := Subscription{}
	values := url.Values{}
	parseSubscriptionParams("create", params, &values)
	err := c.client.post("/customers/"+customerId+"/subscriptions", values, &subscription)
	return &subscription, err
}

// Retrieve loads a subscription.
//
// For more information: https://stripe.com/docs/api#retrieve_subscription
func (c *SubscriptionClient) Retrieve(customerId, id string) (*Subscription, error) {
	subscription := Subscription{}
	err := c.client.get("/customers/"+customerId+"/subscriptions/"+id, nil, &subscription)
	return &subscription, err
}

// Update updates a customers subscription.
//
// For more information: https://stripe.com/docs/api#update_subscription.
func (c *SubscriptionClient) Update(customerId, id string, params *SubscriptionParams) (*Subscription, error) {
	subscription := Subscription{}
	values := url.Values{}
	parseSubscriptionParams("update", params, &values)
	err := c.client.post("/customers/"+customerId+"/subscriptions/"+id, values, &subscription)
	return &subscription, err
}

// Delete cancels a customers subscription.
//
// For more information: https://stripe.com/docs/api#cancel_subscription.
func (c *SubscriptionClient) Delete(customerId, id string, params *SubscriptionParams) (*Subscription, error) {
	subscription := Subscription{}
	values := url.Values{}
	parseSubscriptionParams("cancel", params, &values)
	err := c.client.delete("/customers/"+customerId+"/subscriptions/"+id, values, &subscription)
	return &subscription, err
}

// All lists the first 10 customers. It calls AllWithFilters with a blank Filters
// so all defaults are used.
//
// For more information: https://stripe.com/docs/api#list_customers
func (c *SubscriptionClient) All(customerId string) (*SubscriptionListResponse, error) {
	return c.AllWithFilters(customerId, Filters{})
}

// AllWithFilters takes a Filters and applies all valid filters for the action.
//
// For more information: https://stripe.com/docs/api#list_customers
func (c *SubscriptionClient) AllWithFilters(customerId string, filters Filters) (*SubscriptionListResponse, error) {
	response := SubscriptionListResponse{}
	values := url.Values{}
	addFiltersToValues([]string{"count", "offset"}, filters, &values)
	err := c.client.get("/customers/"+customerId+"/subscriptions", values, &response)
	return &response, err
}

// parseSubscriptionParams takes a string (method), a pointer to
// SubscriptionParams and a pointer to a url.Values, it iterates over everything
// in the SubscriptionParams struct and Adds what is there to the url.Values.
// The first argument, which is the action we are performing ("update" or
// "cancel") determines what values we look for.
func parseSubscriptionParams(method string, params *SubscriptionParams, values *url.Values) {

	if method == "cancel" {
		if params.AtPeriodEnd {
			values.Add("at_period_end", strconv.FormatBool(params.AtPeriodEnd))
		}
		return
	}

	// Use parseCardParams from cards.go to setup the card param
	if params.CardParams != nil {
		parseCardParams(params.CardParams, values, true)
	}

	addParamsToValues(params, values)
}
