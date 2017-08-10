// Package rebrandly General Overview
// ===================================
//
// The package implement most of the API by rebrandly.com.
// The implementation provide full control over links, by doing CRUD operations,
// listing them, and counting them.
//
// The library contact rebrandly by generating a `Request` struct tat provides
// connection and body information regarding the connection.
//
// There are several helper functions to initialize the `Request` struct by
// adding the proper configuration inside the `Request` struct.
//
// When an operation of the API need to send JSON body, there are two helper
// functions implementation:
//   * InitxxxxEx
//   * Initxxxx
//
// The Ex suffix stand for extended, and provide a mean to fully control the
// JSON that is going to be sent.
//
// The counter part function, takes only minimal parameters that are usually
// mandatory by the API in order to function, but does not add extra fields,
// focusing only on the task rather then the extra functionality.
//
//
// How errors works
// ----------------
//
// The package control two types of errors:
//    1. Problematically errors, using Go's error type
//    2. REST errors - Errors arising from using the API itself
//
// The 2nd error types, are struct based, that parsed from JSON return.
// They are located under the `error_types.go` file, and documented to explain
// what they represent
//
// Before explaining how to use the library, it is important to understand that
// the struct that is returned can be the needed JSON content, or an error
// content.
//
// There is an helper function named IsErrorStruct that takes a struct and check
// to see if it is an error struct, to be able to identify it easily.
//
//
// Basic Usage
// -----------
//
// Here is the most simple means to create a new link.
//   link, err := rebrandl.yInitListLinks(
//      "https://www.youtube.com/watch?v=x53JHab2ng8", "sdd12Wa")
//   if err != nil {
//      panic(err)
//   }
//
//   details, err := link.SendRequest("1234567890")
//
// If everything went well, `details` is now a `LinkRequest` struct that
// hold information regarding the link of https://rebrand.ly/sdd12Wa
// with full details about it.
//
// If there was an error returned by the server, then an error struct will be
// returned.
//
// Any other type of error will be placed on the `err` variable instead.
//
package rebrandly

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
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
	if fields.Tags == nil {
		fields.Tags = []string{}
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

// InitUpdateLinkEx initialize the Request struct with parameters for updating
// an existed link.
//
// Required fields:
//  - Destination
//  - SlashTag
//  - Title
//
// Note: If domain was not provided, then it will change the domain to
// rebrand.ly.
func InitUpdateLinkEx(linkID string, fields LinkRequest) (Request, error) {
	url, err := url.Parse(fmt.Sprintf(requestUpdateLinks, linkID))
	if err != nil {
		return Request{}, err
	}
	if fields.Tags == nil {
		fields.Tags = []string{}
	}

	request := Request{
		Method:     http.MethodPost,
		URL:        *url,
		ActionType: ActionTypeLinkUpdate,
		Operation:  fields,
	}
	return request, nil
}

// InitUpdateLink changes the destination and/or slagTag of an existed link
func InitUpdateLink(linkID, destination, slagTag, title string) (Request, error) {
	return InitUpdateLinkEx(linkID, LinkRequest{
		Destination: destination,
		SlashTag:    slagTag,
		Title:       title,
	})
}

// InitDeleteLink deletes a linkid. It is possible to delete it temporarily
// by setting trash to true
func InitDeleteLink(linkID string, trash bool) (Request, error) {
	url, err := url.Parse(fmt.Sprintf(requestDeleteLink, linkID))
	if err != nil {
		return Request{}, err
	}
	q := url.Query()
	q.Add("trash", strconv.FormatBool(trash))
	url.RawQuery = q.Encode()

	request := Request{
		Method:     http.MethodDelete,
		URL:        *url,
		ActionType: ActionTypeLinkDelete,
		Operation:  nil,
	}
	return request, nil
}

// InitLinkDetails returns information on a linkID
func InitLinkDetails(linkID string) (Request, error) {
	url, err := url.Parse(fmt.Sprintf(requestLinkDetails, linkID))
	if err != nil {
		return Request{}, err
	}

	request := Request{
		Method:     http.MethodGet,
		URL:        *url,
		ActionType: ActionTypeLinkDetails,
		Operation:  nil,
	}

	return request, nil
}

// InitListLinks initialize a request for a list of all links based on
// filters, order and pagination
func InitListLinks(favorite bool, status, domainID string,
	orderPagination OrderPagination) (Request, error) {

	url, err := url.Parse(requestListLinks)
	if err != nil {
		return Request{}, err
	}
	orderAndPaginationURL(url, orderPagination)
	q := url.Query()
	q.Add("favorite", strconv.FormatBool(favorite))
	if status != "" {
		q.Add("status", status)
	}
	if domainID != "" {
		q.Add("domain.id", domainID)
	}
	url.RawQuery = q.Encode()

	request := Request{
		Method:     http.MethodGet,
		URL:        *url,
		ActionType: ActionTypeLinkList,
		Operation:  nil,
	}
	return request, nil
}

// InitLinkCount initialize a request for Counting the number of existed links
func InitLinkCount(favourite bool, status, domain string) (Request, error) {
	url, err := url.Parse(requestLinkCount)
	if err != nil {
		return Request{}, err
	}
	q := url.Query()
	q.Add("favourite", strconv.FormatBool(favourite))
	if status != "" {
		q.Add("status", status)
	}
	if domain != "" {
		q.Add("domain.id", domain)
	}
	url.RawQuery = q.Encode()

	request := Request{
		Method:     http.MethodGet,
		URL:        *url,
		ActionType: ActionTypeLinkCount,
		Operation:  nil,
	}
	return request, nil
}

// InitDomainDetails initialize details regarding a domain id
func InitDomainDetails(domainID string) (Request, error) {
	url, err := url.Parse(fmt.Sprintf(requestDomainDetails, domainID))
	if err != nil {
		return Request{}, err
	}
	request := Request{
		Method:     http.MethodGet,
		URL:        *url,
		ActionType: ActionTypeDomainDetails,
		Operation:  nil,
	}
	return request, nil
}

// InitDomainList initialize the domain list with filters, ordering and
// pagination support
func InitDomainList(active bool, domainType string,
	orderPagination OrderPagination) (Request, error) {

	url, err := url.Parse(requestDomainList)
	if err != nil {
		return Request{}, err
	}
	orderAndPaginationURL(url, orderPagination)

	q := url.Query()
	q.Add("active", strconv.FormatBool(active))
	if domainType != "" {
		q.Add("type", domainType)
	}
	url.RawQuery = q.Encode()

	request := Request{
		Method:     http.MethodGet,
		URL:        *url,
		ActionType: ActionTypeDomainList,
		Operation:  nil,
	}

	return request, nil
}

// InitDomainCount initialize the request for counting the number of domains
// available based on filters
func InitDomainCount(active bool, domainType string) (Request, error) {
	url, err := url.Parse(requestDomainCount)
	if err != nil {
		return Request{}, err
	}

	q := url.Query()
	q.Add("active", strconv.FormatBool(active))
	if domainType != "" {
		q.Add("type", domainType)
	}
	url.RawQuery = q.Encode()

	request := Request{
		Method:     http.MethodGet,
		URL:        *url,
		ActionType: ActionTypeDommainCount,
		Operation:  nil,
	}

	return request, nil
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

	return statusCodeToStruct(r, resp.StatusCode, body)
}
