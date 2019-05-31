package test

import "io"

//Request ... Object that handle a request
type Request struct {
	method string
	target Target
	body   io.Reader
}

//NewRequest ... Creates a new Request to handle Request
//<NewRequest returns a new incoming server Request, suitable for passing to an http.Handler for testing.>
func NewRequest(m, target string, body io.Reader) (Request, error) {

	return Request{}, nil
}
