package stripe

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestErrorMessage(t *testing.T) {
	err := ErrorResponse{}
	err.Err.Message = "Error Message"
	assert.Equal(t, err.Error(), err.Err.Message)
}
