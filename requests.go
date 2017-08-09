package rebrandly

import (
	"encoding/json"
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
	if r.Operation != nil {
		structToJSON, err := json.Marshal(r.Operation)
		if err != nil {
			return nil, err
		}
		_, err = io.ReadFull(reader, structToJSON)
		if err != nil {
			return nil, err
		}
	}
	client := &http.Client{}
	req, err := http.NewRequest(r.Method, r.URL.String(), reader)
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

	return r.statusCodeToStruct(resp.StatusCode, body)
}

func (r Request) statusCodeToStruct(statusCode int, body []byte) (result interface{}, err error) {
	err = json.Unmarshal(body, &result)
	return
}
