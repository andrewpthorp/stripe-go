package stripe

import (
	"net/url"
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

type PlanListResponse struct {
	ListResponse
	Data []Plan `json:"data"`
}

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

// All lists the first 10 plans. It calls AllWithFilters with a blank Filters so
// all defaults are used.
//
// For more information: https://stripe.com/docs/api#list_plans
func (p *PlanClient) All() (*PlanListResponse, error) {
	return p.AllWithFilters(Filters{})
}

// AllWithFilters takes a Filters and applies all valid filters for the action.
//
// For more information: https://stripe.com/docs/api#list_plans
func (p *PlanClient) AllWithFilters(filters Filters) (*PlanListResponse, error) {
	response := PlanListResponse{}
	values := url.Values{}
	addFiltersToValues([]string{"count", "offset"}, filters, &values)
	err := get("/plans", values, &response)
	return &response, err
}

// parsePlanParams takes a pointer to a PlanParams and a pointer to a url.Values,
// it iterates over everything in the PlanParams struct and Adds what is there
// to the url.Values.
func parsePlanParams(params *PlanParams, values *url.Values) {

	// Use parseMetaData from metadata.go to setup the metadata param
	if params.Metadata != nil {
		parseMetadata(params.Metadata, values)
	}

	addParamsToValues(params, values)
}
