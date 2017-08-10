package rebrandly

import "net/url"

const (
	contentType     = "application/json"
	rebrandlyAPIURL = "https://api.rebrandly.com/"
)

const (
	requestCreateLinks   = string(rebrandlyAPIURL + "v1/links")
	requestUpdateLinks   = string(rebrandlyAPIURL + "v1/links/%s")
	requestLinkDetails   = string(rebrandlyAPIURL + "v1/links/%s")
	requestDeleteLink    = string(rebrandlyAPIURL + "v1/links/%s")
	requestListLinks     = string(rebrandlyAPIURL + "v1/links")
	requestLinkCount     = string(rebrandlyAPIURL + "v1/links/count")
	requestDomainDetails = string(rebrandlyAPIURL + "v1/domains/%s")
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

// OrderDirType is an enum string type
type OrderDirType string

// enum for OrderDirType
const (
	OrderDirTypeAsc  OrderDirType = "asc"
	OrderDirTypeDesc OrderDirType = "desc"
	OrderDirTypeNone OrderDirType = ""
)

// OrderPagination holds fields to help create list actions for order
// and for pagination
type OrderPagination struct {
	// Order by field - based on the list type, please see documentation for that
	OrderBy string
	// Order direction - default is usually desc
	OrderDir OrderDirType
	// How many records to skip - default 0
	Offset uint64
	// Limit the number of records - default 100
	Limit uint64
}
