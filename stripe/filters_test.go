package stripe

import (
	"github.com/bmizerany/assert"
	"net/url"
	"testing"
)

func TestAddFiltersToValues(t *testing.T) {
	filters := Filters{
		"used":   "yup",
		"unused": "nope",
	}
	values := url.Values{}
	addFiltersToValues([]string{"used"}, filters, &values)
	assert.Equal(t, values.Get("used"), "yup")
	assert.NotEqual(t, values.Get("unused"), "nope")
}
