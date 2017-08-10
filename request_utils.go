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
		result = badRequest

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
		result = unauthorized

	case http.StatusForbidden:
		var badRequest InvalidFormatResponse
		err = json.Unmarshal(body, &badRequest)
		result = badRequest

	case http.StatusNotFound:
		var notFound NotFoundResponse
		err = json.Unmarshal(body, &notFound)
		result = notFound
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
