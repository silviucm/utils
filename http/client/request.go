package client

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// RequestOptions encapsulates a series of http client connectivity options
// from both an encryption and connectivity perspectives.
type RequestOptions struct {
	InsecureSkipVerify bool // Allow insecure https requests
	ClientTimeout      int  // Request timeout, in seconds
}

// SendRequest connects to the url using a request with the specified type, headers and parameters.
// As soon as the content is read, return it and end the function.
//
// The options argument is not mandatory. If not present, the client and tls options are set as follows:
// InsecureSkipVerify: false
// ClientTimeout: client.DEFAULT_REQUEST_TIMEOUT_SECONDS
func SendRequest(url string, requestType string,
	headers map[string]string,
	parameters map[string]string,
	options ...RequestOptions) ([]byte, error) {

	content, _, err := sendRequest(nil, url, requestType, headers, parameters, options...)
	return content, err
}

// SendRequestToWriter connects to the url using a request with the specified type, headers and parameters.
// The content is then written to the specified io.Writer indicated by destination.
// The method returns the number of bytes written, and an error, if applicable.
//
// The options argument is not mandatory. If not present, the client and tls options are set as follows:
// InsecureSkipVerify: false
// ClientTimeout: client.DEFAULT_REQUEST_TIMEOUT_SECONDS
func SendRequestToWriter(destination io.Writer, url string, requestType string,
	headers map[string]string,
	parameters map[string]string,
	options ...RequestOptions) (int64, error) {

	if destination == nil {
		return 0, fmt.Errorf("SendRequestToWriter: the provided destination io.Writer is nil")
	}

	_, written, err := sendRequest(destination, url, requestType, headers, parameters, options...)
	return written, err
}

func sendRequest(destination io.Writer, url string, requestType string,
	headers map[string]string,
	parameters map[string]string,
	options ...RequestOptions) ([]byte, int64, error) {

	defaultOptions := RequestOptions{
		InsecureSkipVerify: DO_NOT_ALLOW_INSECURE_HTTPS,
		ClientTimeout:      DEFAULT_REQUEST_TIMEOUT_SECONDS,
	}

	// If custom options are passed, override the default options
	if len(options) > 0 {
		defaultOptions.ClientTimeout = options[0].ClientTimeout
		defaultOptions.InsecureSkipVerify = options[0].InsecureSkipVerify
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: defaultOptions.InsecureSkipVerify},
	}

	client := http.Client{
		Timeout:   (time.Duration(defaultOptions.ClientTimeout) * time.Second),
		Transport: tr,
	}

	var request *http.Request
	var resp *http.Response
	var err error

	if strings.ToLower(requestType) == "get" {
		// For "GET", join the parameters to the url, we don't need a request body
		request, err = http.NewRequest(requestType,
			strings.Join([]string{url, ConvertMapToRequestParams(parameters, requestType)}, ""), nil)
		if err != nil {
			return nil, 0, err
		}

	} else {
		requestBody := bytes.NewBufferString(ConvertMapToRequestParams(parameters, requestType))
		request, err = http.NewRequest(requestType, url, requestBody)
		if err != nil {
			return nil, 0, err
		}
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	// prepare the request headers
	for currentHeaderKey, currentHeaderValue := range headers {
		request.Header.Add(currentHeaderKey, currentHeaderValue)
	}

	resp, err = client.Do(request)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	// Scenario: no destination writer provided; try to read all content, and
	//
	if destination == nil {
		// read the body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, 0, err
		}
		return body, int64(len(body)), nil
	}

	written, err := io.Copy(destination, resp.Body)
	if err != nil {
		return nil, written, err
	}
	return nil, written, nil

}

// IsTimeoutError returns true if the provided error argument is a standard library
// http error with the Timeout property set to true.
func IsTimeoutError(err error) bool {
	if err == nil {
		return false
	}
	if httpErr, ok := err.(interface {
		Timeout() bool
	}); ok {
		return httpErr.Timeout()
	}
	return false
}

// ConvertMapToRequestParams accepts a map of strings representing the form parameters,
// and constructs a url-encoded string to be used with both GET or POST http requests.
// For POST requests in particular, it might make sense to set the Content-Type header
// to "application/x-www-form-urlencoded" as below:
//
//	requestBody := bytes.NewBufferString(convertMapToRequestParams(paramsMap, "POST"))
//	req, err := http.NewRequest("POST", requestUrl, requestBody)
//	if err != nil {
//		logError(err)
//		return
//	}
//	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
func ConvertMapToRequestParams(parameters map[string]string, requestType string) string {

	if parameters == nil {
		return ""
	}
	if len(parameters) == 0 {
		return ""
	}

	urlValues := url.Values{}
	for k, v := range parameters {
		urlValues.Add(k, v)
	}

	if strings.ToLower(requestType) == "get" {
		return strings.Join([]string{"?", urlValues.Encode()}, "")
	}
	return urlValues.Encode()
}
