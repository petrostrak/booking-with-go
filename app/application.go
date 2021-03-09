package app

import "net/http"

const (
	portNumber = ":8000"
)

// StartApp func will initializa the url mapping
// and the server
func StartApp() {
	urlMapping()

	if err := http.ListenAndServe(portNumber, nil); err != nil {
		panic(err)
	}
}
