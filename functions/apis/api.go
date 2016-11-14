package main

import (
	"fmt"
	"net/http"
)

// NewHTTPHandler creates a http.Handler for handling HTTP requests.
// The default implementation creates a http.ServerMux. You may replace it with your favorite third-party
// HTTP routing object, such as ozzo-routing, echo, gin.
func NewHTTPHandler() http.Handler {
	// set up the routing using the standard ServerMux or third-party HTTP router
	mux := http.NewServeMux()
	mux.HandleFunc("/foo", foo)
	mux.HandleFunc("/bar", bar)
	return mux
}

// foo handles the "/foo" request.
func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello")
}

// bar handles the "/bar" request.
func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%v %v", req.Method, req.URL)
}
