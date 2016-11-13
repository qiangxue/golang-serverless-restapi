package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/apex/go-apex"
)

func main() {
	// set up the routing using the standard ServerMux or third-party HTTP router
	mux := http.NewServeMux()
	mux.HandleFunc("/foo", foo)
	mux.HandleFunc("/bar", bar)

	// register the handler to handle Lambda events
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		req, err := ParseRequest(event)
		if err != nil {
			return FormatError(http.StatusBadRequest, err), nil
		}

		res := httptest.NewRecorder()

		// use ServerMux or third-party HTTP router to handle the request
		mux.ServeHTTP(res, req)

		return FormatResponse(res), nil
	})
}

// foo is an HTTP handler for "/foo"
func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello")
}

// bar is an HTTP handler for "/bar"
func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%v %v", req.Method, req.URL)
}
