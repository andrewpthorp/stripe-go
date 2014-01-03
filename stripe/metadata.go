package stripe

import (
	"net/url"
)

// parseMetadata takes a pointer to a map of strings to strings and a pointer to
// a map of url.Values. It iterates over each element in the map and adds it to
// the url.Values.
func parseMetadata(metadata map[string]string, values *url.Values) {
	for k, v := range metadata {
		values.Add("metadata["+k+"]", v)
	}
}
