package rebrandly

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// InitCreateLinkEx initialize the Request struct with parameters for creating
// a link.
// The function uses LinkRequest struct to better control the creation of a new
// link.
//
// IMPORTANT:
//  - Empty slagtag means that rebrandly create custom link.
//  - If no domain is configured, then `rebrand.ly` is used by them by default.
func InitCreateLinkEx(fields LinkRequest) (Request, error) {
	url, err := url.Parse(requestCreateLinks)
	if err != nil {
		return Request{}, err
	}
	request := Request{
		Method:     http.MethodPost,
		URL:        *url,
		ActionType: ActionTypeLinkCreate,
		Operation:  fields,
	}
	return request, nil
}

// InitCreateLink initialize the Request struct with parameters for creating
// a link.
// The initialization is only with mandatory fileds.
// For advanced initialization, use the InitCreateLinkEx func instead
func InitCreateLink(destination string, slagTag string) (Request, error) {
	return InitCreateLinkEx(LinkRequest{
		Destination: destination,
		SlashTag:    slagTag,
	})
}

// SendRequest send a request to rebrandly.
// If everything goes well, the return is the answer by the HTTP request
// If there was internal issue, an error return
func (r Request) SendRequest(apiKey string) (interface{}, error) {
	var reader io.Reader
	var structToJSON []byte
	var err error
	if r.Operation != nil {
		structToJSON, err = json.Marshal(r.Operation)
		if err != nil {
			return nil, err
		}

		reader = bytes.NewReader(structToJSON)
	}
	client := &http.Client{}
	var req *http.Request
	if len(structToJSON) > 0 {
		req, err = http.NewRequest(r.Method, r.URL.String(), reader)
	} else {
		req, err = http.NewRequest(r.Method, r.URL.String(), nil)
	}
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("apikey", apiKey)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Printf("StatusCode: %d, body: %#v\n", resp.StatusCode, string(body))
	return r.statusCodeToStruct(resp.StatusCode, body)
}

func (r Request) statusCodeToStruct(statusCode int, body []byte) (result interface{}, err error) {
	switch statusCode {
	case http.StatusOK:
	case http.StatusBadRequest:
	case http.StatusUnauthorized:
		var unauthorized UnauthorizedResponse
		err = json.Unmarshal(body, &unauthorized)
		result = unauthorized
	case http.StatusForbidden:
		var badRequest InvalidFormatResponse
		err = json.Unmarshal(body, &badRequest)
		result = badRequest
	case http.StatusNotFound:
	case http.StatusInternalServerError,
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout:
	default:
		return nil, fmt.Errorf("Unsupported StatusCode: %d", statusCode)
	}
	// err = json.Unmarshal(body, &result)
	return
}
