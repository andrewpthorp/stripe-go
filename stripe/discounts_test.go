package stripe

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestDiscountsDelete(t *testing.T) {
	setup()
	defer teardown()
	handleWithJSON("/customers/id/discount", "delete.json")
	res, _ := client.Discounts.Delete("id")
	assert.Equal(t, res.Deleted, true)
}
