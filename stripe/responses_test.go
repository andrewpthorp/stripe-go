package stripe

import (
  "testing"
  "github.com/bmizerany/assert"
)

func TestErrorMessage(t *testing.T){
  err := ErrorResponse{}
  err.Err.Message = "Error Message"
  assert.Equal(t, err.Error(), err.Err.Message)
}
