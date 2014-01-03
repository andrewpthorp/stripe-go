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

type SubscriptionClient struct{}

// Update updates a customers subscription.
//
// For more information: https://stripe.com/docs/api#update_subscription.
func (c *SubscriptionClient) Update(customerId string, params *SubscriptionParams) (*Subscription, error) {
	subscription := Subscription{}
	values := url.Values{}
	parseSubscriptionParams("update", params, &values)
	err := post("/customers/"+customerId+"/subscription", values, &subscription)
	return &subscription, err
}

// Cancel cancels a customers subscription.
//
// For more information: https://stripe.com/docs/api#cancel_subscription.
func (c *SubscriptionClient) Cancel(customerId string, params *SubscriptionParams) (*Subscription, error) {
	subscription := Subscription{}
	values := url.Values{}
	parseSubscriptionParams("cancel", params, &values)
	err := delete("/customers/"+customerId+"/subscription", values, &subscription)
	return &subscription, err
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

	if params.Plan != "" {
		values.Add("plan", params.Plan)
	}

	if params.Coupon != "" {
		values.Add("coupon", params.Coupon)
	}

	// TODO: What to do with Prorate?

	if params.TrialEnd != 0 {
		values.Add("trial_end", strconv.Itoa(params.TrialEnd))
	}

	if params.Quantity != 0 {
		values.Add("quantity", strconv.Itoa(params.Quantity))
	}

	if params.ApplicationFeePercent != 0.0 {
		values.Add("application_fee_percent", strconv.FormatFloat(params.ApplicationFeePercent, 'f', 2, 32))
	}

	if params.AtPeriodEnd {
		values.Add("at_period_end", strconv.FormatBool(params.AtPeriodEnd))
	}
}
