package rebrandly

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func statusCodeToStruct(r Request, statusCode int, body []byte) (result interface{}, err error) {
	switch statusCode {
	case http.StatusOK:
		switch r.ActionType {
		case ActionTypeDomainList:
			var domainList DomainRequestList
			err = json.Unmarshal(body, &domainList)
			result = domainList

		case ActionTypeDomainDetails:
			var domain DomainRequest
			err = json.Unmarshal(body, &domain)
			result = domain

		case ActionTypeLinkCount,
			ActionTypeDommainCount:
			var linkCount CountRequest
			err = json.Unmarshal(body, &linkCount)
			result = linkCount

		case ActionTypeLinkList:
			var linkList LinkRequestList
			err = json.Unmarshal(body, &linkList)
			result = linkList

		case ActionTypeLinkCreate,
			ActionTypeLinkUpdate,
			ActionTypeLinkDelete,
			ActionTypeLinkDetails:
			var linkRequest LinkRequest
			err = json.Unmarshal(body, &linkRequest)
			result = linkRequest
		}

	case http.StatusBadRequest:
		var badRequest BadRequestResponse
		err = json.Unmarshal(body, &badRequest)
		if err == nil {
			err = badRequest
		}

	case http.StatusUnauthorized:
		var unauthorized UnauthorizedResponse
		if string(body) == "Unauthorized" {
			unauthorized = UnauthorizedResponse{
				Message: string(body),
				Code:    ErrorCodeUnauthorized,
			}
		} else {
			err = json.Unmarshal(body, &unauthorized)
		}
		if err == nil {
			err = unauthorized
		}

	case http.StatusForbidden:
		var badRequest InvalidFormatResponse
		err = json.Unmarshal(body, &badRequest)
		if err == nil {
			err = badRequest
		}

	case http.StatusNotFound:
		var notFound NotFoundResponse
		err = json.Unmarshal(body, &notFound)
		if err == nil {
			err = notFound
		}
	case http.StatusInternalServerError,
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout:

		var serverErr ServerErrorResponse
		err = json.Unmarshal(body, &serverErr)
		if err == nil {
			err = serverErr
		}

	default:
		return nil, fmt.Errorf("Unsupported StatusCode: %d", statusCode)
	}
	return
}

func orderAndPaginationURL(u *url.URL, orderPagination OrderPagination) {
	q := u.Query()
	if orderPagination.OrderBy != "" {
		q.Add("orderBy", orderPagination.OrderBy)
	}
	if orderPagination.OrderDir != OrderDirTypeNone {
		q.Add("orderDir", string(orderPagination.OrderDir))
	}
	if orderPagination.Limit > 0 {
		q.Add("limit", strconv.FormatUint(orderPagination.Limit, 10))
	}
	q.Add("offset", strconv.FormatUint(orderPagination.Offset, 10))
	u.RawQuery = q.Encode()
}
