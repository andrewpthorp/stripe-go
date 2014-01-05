package stripe

import (
	"net/url"
)

// TODO: There is probably a better way to do this.
type EventData struct {
	Object map[string]interface{} `json:"object"`
}

type Event struct {
	Id              string     `json:"id"`
	Object          string     `json:"object"`
	Data            *EventData `json:"data"`
	Livemode        bool       `json:"livemode"`
	Created         int64      `json:"created"`
	PendingWebhooks int64      `json:"pending_webhooks"`
	Type            string     `json:"type"`
	Request         string     `json:"request"`
}

type EventListResponse struct {
	Object string   `json:"object"`
	Url    string   `json:"url"`
	Count  int      `json:"count"`
	Data   []*Event `json:"data"`
}

type EventClient struct{}

// Retrieve loads a event.
//
// For more information: https://stripe.com/docs/api#retrieve_event
func (p *EventClient) Retrieve(id string) (*Event, error) {
	event := Event{}
	err := get("/events/"+id, nil, &event)
	return &event, err
}

// All lists the first 10 events. It calls AllWithFilters with a blank Filters
// so all defaults are used.
//
// For more information: https://stripe.com/docs/api#list_events
func (c *EventClient) All() (*EventListResponse, error) {
	return c.AllWithFilters(Filters{})
}

// AllWithFilters takes a Filters and applies all valid filters for the action.
//
// For more information: https://stripe.com/docs/api#list_events
func (c *EventClient) AllWithFilters(filters Filters) (*EventListResponse, error) {
	response := EventListResponse{}
	values := url.Values{}
	addFiltersToValues([]string{"count", "offset", "type"}, filters, &values)
	err := get("/events", values, &response)
	return &response, err
}
