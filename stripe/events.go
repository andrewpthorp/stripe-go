package stripe

import (
	"net/url"
	"strconv"
)

// TODO: There is probably a better way to do this.
type EventData struct {
	Object map[string]interface{} `json:"object"`
}

type Event struct {
	Id              string    `json:"id"`
	Object          string    `json:"object"`
	Data            EventData `json:"data"`
	Livemode        bool      `json:"livemode"`
	Created         int64     `json:"created"`
	PendingWebhooks int64     `json:"pending_webhooks"`
	Type            string    `json:"type"`
	Request         string    `json:"request"`
}

// The EventClient is the receiver for most standard event related endpoints.
type EventClient struct{}

// Retrieve loads a event.
//
// For more information: https://stripe.com/docs/api#retrieve_event
func (p *EventClient) Retrieve(id string) (*Event, error) {
	event := Event{}
	err := get("/events/"+id, nil, &event)
	return &event, err
}

// List lists the first 10 events. It calls ListCount with 10 as the count and
// 0 as the offset, which are the defaults in the Stripe API.
//
// For more information: https://stripe.com/docs/api#list_events
func (c *EventClient) List() ([]*Event, error) {
	return c.ListCount(10, 0)
}

// ListCount lists `count` events starting at `offset`.
//
// For more information: https://stripe.com/docs/api#list_events
func (c *EventClient) ListCount(count, offset int) ([]*Event, error) {
	type events struct{ Data []*Event }
	list := events{}

	params := url.Values{
		"count":  {strconv.Itoa(count)},
		"offset": {strconv.Itoa(offset)},
	}

	err := get("/events", params, &list)
	return list.Data, err
}
