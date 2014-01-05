package stripe

import (
  "net/url"
  "testing"
  "github.com/bmizerany/assert"
)

func TestAddFiltersToValues(t *testing.T){
  filters := Filters{
    "used": "yup",
    "unused": "nope",
  }
  values := url.Values{}
  addFiltersToValues([]string{"used"}, filters, &values)
  assert.Equal(t, values.Get("used"), "yup")
  assert.NotEqual(t, values.Get("unused"), "nope")
}
