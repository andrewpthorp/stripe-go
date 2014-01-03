package stripe

import (
	"net/url"
	"strconv"
)

type Plan struct {
	Id              string   `json:"id"`
	Object          string   `json:"object"`
	Livemode        bool     `json:"livemode"`
	Amount          int64    `json:"amount"`
	Currency        string   `json:"currency"`
	Interval        string   `json:"interval"`
	IntervalCount   int64    `json:"interval_count"`
	Name            string   `json:"name"`
	TrialPeriodDays int64    `json:"trial_period_days"`
	Metadata        Metadata `json:"metadata"`
}

// Delete deletes a plan.
//
// For more information: https://stripe.com/docs/api#delete_plan
func (p *Plan) Delete() (*DeleteResponse, error) {
	response := DeleteResponse{}
	err := delete("/plans/"+p.Id, nil, &response)
	return &response, err
}

// Update updates a plan.
//
// For more information: https://stripe.com/docs/api#update_plan
func (p *Plan) Update(params *PlanParams) (*Plan, error) {
	values := url.Values{}
	parsePlanParams(params, &values)
	err := post("/plans/"+p.Id, values, p)
	return p, err
}

// The PlanClient is the receiver for most standard plan related endpoints.
type PlanClient struct{}

// Create creates a plan.
//
// For more information: https://stripe.com/docs/api#create_plan
func (p *PlanClient) Create(params *PlanParams) (*Plan, error) {
	plan := Plan{}
	values := url.Values{}
	parsePlanParams(params, &values)
	err := post("/plans", values, &plan)
	return &plan, err
}

// Retrieve loads a plan.
//
// For more information: https://stripe.com/docs/api#retrieve_plan
func (p *PlanClient) Retrieve(id string) (*Plan, error) {
	plan := Plan{}
	err := get("/plans/"+id, nil, &plan)
	return &plan, err
}

// Update updates a plan.
//
// For more information: https://stripe.com/docs/api#update_plan
func (p *PlanClient) Update(id string, params *PlanParams) (*Plan, error) {
	plan := Plan{}
	values := url.Values{}
	parsePlanParams(params, &values)
	err := post("/plans/"+id, values, &plan)
	return &plan, err
}

// Delete deletes a plan.
//
// For more information: https://stripe.com/docs/api#delete_plan
func (p *PlanClient) Delete(id string) (*DeleteResponse, error) {
	response := DeleteResponse{}
	err := delete("/plans/"+id, nil, &response)
	return &response, err
}

// List lists the first 10 plans. It calls ListCount with 10 as
// the count and 0 as the offset, which are the defaults in the Stripe API.
//
// For more information: https://stripe.com/docs/api#list_plans
func (p *PlanClient) List() ([]*Plan, error) {
	return p.ListCount(10, 0)
}

// ListCount lists `count` plans starting at `offset`.
//
// For more information: https://stripe.com/docs/api#list_plans
func (p *PlanClient) ListCount(count, offset int) ([]*Plan, error) {
	type plans struct{ Data []*Plan }
	list := plans{}

	params := url.Values{
		"count":  {strconv.Itoa(count)},
		"offset": {strconv.Itoa(offset)},
	}

	err := get("/plans", params, &list)
	return list.Data, err
}

// parsePlanParams takes a pointer to a PlanParams and a pointer to a url.Values,
// it iterates over everything in the PlanParams struct and Adds what is there
// to the url.Values.
func parsePlanParams(params *PlanParams, values *url.Values) {

	// Use parseMetaData from metadata.go to setup the metadata param
	if params.Metadata != nil {
		parseMetadata(params.Metadata, values)
	}

	if params.Id != "" {
		values.Add("id", params.Id)
	}

	if params.Amount != 0 {
		values.Add("amount", strconv.Itoa(params.Amount))
	}

	if params.Currency != "" {
		values.Add("currency", params.Currency)
	}

	if params.Interval != "" {
		values.Add("interval", params.Interval)
	}

	if params.IntervalCount != 0 {
		values.Add("interval_count", strconv.Itoa(params.IntervalCount))
	}

	if params.Name != "" {
		values.Add("name", params.Name)
	}

	if params.TrialPeriodDays != 0 {
		values.Add("trial_period_days", strconv.Itoa(params.TrialPeriodDays))
	}
}
