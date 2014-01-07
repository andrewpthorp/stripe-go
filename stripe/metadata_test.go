package stripe

import (
	"github.com/bmizerany/assert"
	"net/url"
	"testing"
)

func TestParseMetadata(t *testing.T) {
	meta := Metadata{
		"foo": "bar",
	}
	values := url.Values{}
	parseMetadata(meta, &values)
	assert.Equal(t, values.Get("metadata[foo]"), "bar")
}
