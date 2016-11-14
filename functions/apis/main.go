package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/apex/go-apex"
)

func main() {
	// set up the HTTP routing
	handler := NewHTTPHandler()

	// register the Lambda event handler
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		req, err := ParseRequest(event)
		if err != nil {
			return FormatError(http.StatusBadRequest, err), nil
		}

		res := httptest.NewRecorder()

		// handle the HTTP request
		handler.ServeHTTP(res, req)

		return FormatResponse(res), nil
	})
}
