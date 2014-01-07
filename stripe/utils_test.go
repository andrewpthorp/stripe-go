package stripe

import (
	"github.com/bmizerany/assert"
	"net/url"
	"testing"
)

type Mock struct {
	MockString string  `stripe_field:"string_value"`
	MockBool   bool    `stripe_field:"bool_value"`
	MockInt    int     `stripe_field:"int_value"`
	MockFloat  float64 `stripe_field:"float64_value"`
}

func TestAddParamsToValues(t *testing.T) {
	mock := Mock{
		MockString: "foo",
		MockBool:   true,
		MockInt:    10,
		MockFloat:  25.50,
	}
	values := url.Values{}
	addParamsToValues(&mock, &values)
	assert.Equal(t, values.Get("string_value"), "foo")
	assert.Equal(t, values.Get("bool_value"), "true")
	assert.Equal(t, values.Get("int_value"), "10")
	assert.Equal(t, values.Get("float64_value"), "25.50")
}

func TestAttributes(t *testing.T) {
	mock := Mock{
		MockString: "foo",
		MockBool:   true,
		MockInt:    10,
		MockFloat:  25.50,
	}
	attrs := attributes(&mock)
	assert.Equal(t, attrs["MockString"].Name(), "string")
	assert.Equal(t, attrs["MockBool"].Name(), "bool")
	assert.Equal(t, attrs["MockInt"].Name(), "int")
	assert.Equal(t, attrs["MockFloat"].Name(), "float64")
}

func TestGetTag(t *testing.T) {
	tag := getTag(new(Mock), "stripe_field", "MockString")
	assert.Equal(t, tag, "string_value")
}

func TestGetString(t *testing.T) {
	var mock Mock

	// Zero value
	mock = Mock{}
	assert.Equal(t, getString(&mock, "MockString"), "")

	// Non zero value
	mock = Mock{MockString: "foo"}
	assert.Equal(t, getString(&mock, "MockString"), "foo")
}

func TestGetBool(t *testing.T) {
	var mock Mock

	// Zero value
	mock = Mock{}
	assert.Equal(t, getBool(&mock, "MockBool"), "")

	// Non zero value
	mock = Mock{MockBool: true}
	assert.Equal(t, getBool(&mock, "MockBool"), "true")
}

func TestGetInt(t *testing.T) {
	var mock Mock

	// Zero value
	mock = Mock{}
	assert.Equal(t, getInt(&mock, "MockInt"), "")

	// Non zero value
	mock = Mock{MockInt: 10}
	assert.Equal(t, getInt(&mock, "MockInt"), "10")
}

func TestGetFloat64(t *testing.T) {
	var mock Mock

	// Zero value
	mock = Mock{}
	assert.Equal(t, getFloat64(&mock, "MockFloat"), "")

	// Non zero value
	mock = Mock{MockFloat: 20.5}
	assert.Equal(t, getFloat64(&mock, "MockFloat"), "20.50")
}

func TestGetField(t *testing.T) {
	mock := Mock{MockString: "foobar"}
	assert.Equal(t, getField(&mock, "MockString").String(), "foobar")
}
