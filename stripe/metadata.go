package stripe

import (
	"net/url"
)

// Metadata is a map of strings to strings
type Metadata map[string]string

// parseMetadata takes a pointer to a map of strings to strings and a pointer to
// a map of url.Values. It iterates over each element in the map and adds it to
// the url.Values.
func parseMetadata(metadata Metadata, values *url.Values) {
	for k, v := range metadata {
		values.Add("metadata["+k+"]", v)
	}
}
