package stripe

import (
  "fmt"
  "io/ioutil"
	"net/http"
	"net/http/httptest"
  "os"
  "path"
)

var (
	client   Client
	serveMux *http.ServeMux
	server   *httptest.Server
)

// setup starts a new Server, using a ServeMux, it also initializes a client
// with the url of the new Server. The typical use for this will be:
//
//     setup()
//     defer teardown()
//
// This makes sure the server is started and stopped inside of a test.
func setup() {
	serveMux = http.NewServeMux()
	server = httptest.NewServer(serveMux)
  client = NewClientWith(server.URL, "sk_abc123")
}

// teardown closes the server that is initialized in setup()
func teardown() {
  server.Close()
}

// loadFixture takes a path to a fixture file and returns a string of what the
// file contains.
func loadFixture(f string) string {
  wd, _ := os.Getwd()
  p := path.Join(wd, "..", "fixtures", f)
  c, _ := ioutil.ReadFile(p)
  return string(c)
}

// handleWithJson takes a path and a jason filename, it uses the serveMux to
// handle that path and respond with the json.
func handleWithJSON(path, filename string) {
  serveMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, loadFixture(filename))
  })
}
