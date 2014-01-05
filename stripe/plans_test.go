package stripe

import (
  "strconv"
  "testing"
  "net/url"
  "github.com/bmizerany/assert"
)

func TestPlanCreate(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/plans", "plans/plan.json")
  params := PlanParams{}
  plan, _ := client.Plans.Create(&params)
  assert.Equal(t, plan.Id, "plan_code")
}

func TestPlansRetrieve(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/plans/id", "plans/plan.json")
  plan, _ := client.Plans.Retrieve("id")
  assert.Equal(t, plan.Id, "plan_code")
}

func TestPlansUpdate(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/plans/id", "plans/plan.json")
  plan, _ := client.Plans.Update("id", new(PlanParams))
  assert.Equal(t, plan.Id, "plan_code")
}

func TestPlansDelete(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/plans/id", "delete.json")
  res, _ := client.Plans.Delete("id")
  assert.Equal(t, res.Deleted, true)
}

func TestPlansAll(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/plans", "plans/plans.json")
  plans, _ := client.Plans.All()
  assert.Equal(t, plans.Count, 1)
  assert.Equal(t, plans.Data[0].Id, "plan_code")
}


func TestPlansAllWithFilters(t *testing.T) {
  setup()
  defer teardown()
  handleWithJSON("/plans", "plans/plans.json")
  plans, _ := client.Plans.AllWithFilters(Filters{})
  assert.Equal(t, plans.Count, 1)
  assert.Equal(t, plans.Data[0].Id, "plan_code")
}

func TestParsePlanParams(t *testing.T) {
  params := PlanParams{
    Id: "plan_id",
    Amount: 1000,
    Currency: "USD",
    Interval: "monthly",
    IntervalCount: 2,
    Name: "Plan",
    TrialPeriodDays: 1,
    Metadata: Metadata{
      "foo": "bar",
    },
  }
  values := url.Values{}
  parsePlanParams(&params, &values)
  assert.Equal(t, values.Get("id"), params.Id)
  assert.Equal(t, values.Get("amount"), strconv.Itoa(params.Amount))
  assert.Equal(t, values.Get("currency"), params.Currency)
  assert.Equal(t, values.Get("interval"), params.Interval)
  assert.Equal(t, values.Get("interval_count"), strconv.Itoa(params.IntervalCount))
  assert.Equal(t, values.Get("name"), params.Name)
  assert.Equal(t, values.Get("trial_period_days"), strconv.Itoa(params.TrialPeriodDays))
  assert.Equal(t, values.Get("metadata[foo]"), params.Metadata["foo"])
}
