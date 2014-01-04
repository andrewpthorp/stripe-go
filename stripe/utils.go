package stripe

import (
	"net/url"
	"reflect"
	"strconv"
)

// addParamsToValues takes an interface (usually *SomeTypeParams) and a pointer
// to a url.Values. It iterates over each field in the interface (using
// the attributes method), and adds the value of each field to the url.Values.
func addParamsToValues(params interface{}, values *url.Values) {
	var val string

	for name, mtype := range attributes(params) {
		switch mtype.Name() {
		case "string":
			val = getString(params, name)
		case "int":
			val = getInt(params, name)
		case "float64":
			val = getFloat64(params, name)
		case "bool":
			val = getBool(params, name)
		}

		if val != "" {
			values.Add(getTag(params, "stripe_field", name), val)
		}
	}
}

// attributes takes a struct m and returns a map of strings (field names) to
// reflect.Types (field types).
func attributes(m interface{}) map[string]reflect.Type {
	typ := reflect.TypeOf(m)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	attrs := make(map[string]reflect.Type)
	for i := 0; i < typ.NumField(); i++ {
		p := typ.Field(i)
		if !p.Anonymous {
			attrs[p.Name] = p.Type
		}
	}

	return attrs
}

// getTag gets the tagName tag from an element in a struct. It takes a struct
// and a fieldName (string), and returns the tag named tagName ("stripe_field",
// etc) for that fieldName.
func getTag(m interface{}, tagName, fieldName string) string {
	f, _ := reflect.TypeOf(m).Elem().FieldByName(fieldName)
	return f.Tag.Get(tagName)
}

// getString gets the value of fieldName in the struct m, and returns it.
func getString(m interface{}, fieldName string) string {
	return getField(m, fieldName).String()
}

// getBool gets the value of fieldName in the struct m (bool), converts it to
// a string, and returns the result. If the value is the zero value (false), it
// returns a blank string.
func getBool(m interface{}, fieldName string) string {
	val := getField(m, fieldName).Bool()

	if val {
		opposite, _ := strconv.ParseBool(getTag(m, "opposite", fieldName))

		if opposite {
			return strconv.FormatBool(!val)
		} else {
			return strconv.FormatBool(val)
		}

	} else {
		return ""
	}
}

// getInt gets the value of fieldName in the struct m (int), converts it to a
// string, and returns the result. If the value is the zero value (0), it
// returns a blank string.
func getInt(m interface{}, fieldName string) string {
	val := int(getField(m, fieldName).Int())

	if val == 0 {
		return ""
	} else {
		return strconv.Itoa(val)
	}
}

// getFloat gets the value of fieldName in the struct m (float64), converts it
// to a string, and returns the result. If the value is the zero value (0.0),
// it returns a blank string.
func getFloat64(m interface{}, fieldName string) string {
	val := getField(m, fieldName).Float()

	if val == 0.0 {
		return ""
	} else {
		return strconv.FormatFloat(val, 'f', 2, 32)
	}
}

// getField gets the reflect.Value of fieldName in the struct m.
func getField(m interface{}, fieldName string) reflect.Value {
	val := reflect.ValueOf(m)
	return reflect.Indirect(val).FieldByName(fieldName)
}
