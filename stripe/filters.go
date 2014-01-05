package stripe

import "net/url"

type Filters map[string]string

// addFiltersToValues takes a slice of strings, a Filters, and a pointer to a
// url.Values. If those strings are present in the Filters, it adds them to the
// url.Values.
func addFiltersToValues(keys []string, filters Filters, values *url.Values) {

	for i := 0; i < len(keys); i++ {
		key := keys[i]
		if filters[key] != "" {
			values.Add(key, filters[key])
		}
	}

}
