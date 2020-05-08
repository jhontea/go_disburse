package apicall

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

var (
	// ErrNewRequest :nodoc:
	ErrNewRequest = errors.New("Error create new request")

	// ErrDoRequest :nodoc:
	ErrDoRequest = errors.New("Error when doing request")

	// ErrReadBody :nodoc:
	ErrReadBody = errors.New("Error read response body")

	// TODO: move to env

	// BasicAuthUsername :nodoc:
	BasicAuthUsername = "HyzioY7LP6ZoO7nTYKbG8O4ISkyWnX1JvAEVAhtWKZumooCzqp41"
)

// APICall :nodoc:
type APICall struct {
	URL       string
	Method    string
	FormParam string
	Header    map[string]string
}

// URLHttpResponse :nodoc:
type URLHttpResponse struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"status_code"`
	Body       string      `json:"body"`
	Header     http.Header `json:"header"`
}

// Call call to third party endpoint
func (apicall *APICall) Call() (URLHttpResponse, error) {
	var result URLHttpResponse

	// Http request
	client := &http.Client{}
	req, err := http.NewRequest(apicall.Method, apicall.URL, bytes.NewBuffer([]byte(apicall.FormParam)))
	if err != nil {
		return result, ErrNewRequest
	}

	// Set header
	// -- Content type
	req.Header.Add("Content-Type", "application/json")
	for index, value := range apicall.Header {
		req.Header.Add(index, value)
	}

	req.SetBasicAuth(BasicAuthUsername, "")

	// Do request
	resp, err := client.Do(req)
	if err != nil {
		return result, ErrDoRequest
	}

	defer resp.Body.Close()

	// Get string body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, ErrReadBody
	}

	result.Status = resp.Status
	result.StatusCode = resp.StatusCode
	result.Body = string(body)
	result.Header = resp.Header

	return result, nil
}
