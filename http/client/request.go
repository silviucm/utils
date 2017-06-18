package client

import (
	"bytes"
	"crypto/tls"
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
			return nil, err
		}

	} else {
		requestBody := bytes.NewBufferString(ConvertMapToRequestParams(parameters, requestType))
		request, err = http.NewRequest(requestType, url, requestBody)
		if err != nil {
			return nil, err
		}
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	// prepare the request headers
	for currentHeaderKey, currentHeaderValue := range headers {
		request.Header.Add(currentHeaderKey, currentHeaderValue)
	}

	resp, err = client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read the body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
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
