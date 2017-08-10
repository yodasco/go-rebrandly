package rebrandly

import "net/url"

const (
	contentType     = "application/json"
	rebrandlyAPIURL = "https://api.rebrandly.com/"
)

const (
	requestCreateLinks = string(rebrandlyAPIURL + "v1/links")
	requestUpdateLinks = string(rebrandlyAPIURL + "v1/links/%s")
)

// ActionTypes is an enum of action types
type ActionTypes string

// An enum of all supported actions by this library
const (
	ActionTypeLinkCreate    ActionTypes = "linkcreate"
	ActionTypeLinkUpdate    ActionTypes = "linkupdate"
	ActionTypeLinkDelete    ActionTypes = "linkdelete"
	ActionTypeLinkDetails   ActionTypes = "linkdetails"
	ActionTypeLinkList      ActionTypes = "linklist"
	ActionTypeLinkCount     ActionTypes = "linkcount"
	ActionTypeDomainDetails ActionTypes = "domaindetails"
	ActionTypeDomainList    ActionTypes = "domainlist"
	ActionTypeDommainCount  ActionTypes = "domaincount"
)

// Request is a struct that represent an HTTP request
type Request struct {
	// GET, POST and DELETE
	Method string
	// request Path, and other GET parameters (for pagination, and sorting)
	URL url.URL
	// The type of action to do
	ActionType ActionTypes
	// The struct for the operation to be made
	Operation interface{}
}
