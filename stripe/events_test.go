package stripe

import (
  "testing"
  "github.com/bmizerany/assert"
)

func TestEventsRetrieve(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/events/evt_123456789", "events/event.json")
  event, _ := client.Events.Retrieve("evt_123456789")
  assert.Equal(t, event.Id, "evt_123456789")
  assert.Equal(t, event.Data.Object["plan"].(map[string]interface{})["name"], "Plan")
  assert.Equal(t, event.Data.PreviousAttributes["plan"].(map[string]interface{})["name"], "Monthly Plan")
}

func TestEventsAll(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/events", "events/events.json")
  events, _ := client.Events.All()
  assert.Equal(t, events.Count, 1)
  assert.Equal(t, events.Data[0].Id, "evt_123456789")
}


func TestEventsAllWithFilters(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/events", "events/events.json")
  events, _ := client.Events.AllWithFilters(Filters{})
  assert.Equal(t, events.Count, 1)
  assert.Equal(t, events.Data[0].Id, "evt_123456789")
}
