package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
)

type (
	// LambdaInput is the Lambda input given by AWS API Gateway
	LambdaInput struct {
		Body    string            `json:"body"`
		Headers map[string]string `json:"headers"`
		Method  string            `json:"httpMethod"`
		Path    string            `json:"path"`
		Params  map[string]string `json:"queryStringParameters"`
	}

	// LambdaOutput is the Lambda return result expected by AWS API Gateway
	LambdaOutput struct {
		StatusCode int               `json:"statusCode"`
		Headers    map[string]string `json:"headers"`
		Body       string            `json:"body"`
	}
)

// ParseRequest parses the Lambda event into a standard HTTP request.
func ParseRequest(event json.RawMessage) (*http.Request, error) {
	var input LambdaInput
	if err := json.Unmarshal(event, &input); err != nil {
		return nil, err
	}

	// gather query parameters and HTTP body
	v := url.Values{}
	for key, value := range input.Params {
		v.Set(key, value)
	}
	req, err := http.NewRequest(input.Method, input.Path+"?"+v.Encode(), bytes.NewBufferString(input.Body))
	if err != nil {
		return nil, err
	}

	// gather HTTP headers
	for key, value := range input.Headers {
		req.Header.Set(key, value)
	}
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

// FormatResponse formats a ResponseRecorder into a LambdaOutput.
func FormatResponse(res *httptest.ResponseRecorder) LambdaOutput {
	result := LambdaOutput{
		StatusCode: res.Code,
		Body:       res.Body.String(),
		Headers:    map[string]string{},
	}
	for key := range res.HeaderMap {
		result.Headers[key] = res.HeaderMap.Get(key)
	}
	return result
}

// FormatError formats an error and an HTTP status code into a LambdaOutput.
func FormatError(status int, err error) LambdaOutput {
	bs, _ := json.Marshal(err)
	return LambdaOutput{
		StatusCode: status,
		Body:       string(bs),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}
