package stripe

import (
  "net/url"
  "testing"
  "github.com/bmizerany/assert"
)

func TestParseMetadata(t *testing.T){
  meta := Metadata{
    "foo": "bar",
  }
  values := url.Values{}
  parseMetadata(meta, &values)
  assert.Equal(t, values.Get("metadata[foo]"), "bar")
}
