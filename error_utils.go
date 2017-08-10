package rebrandly

import (
	"reflect"
)

// IsErrorStruct is an helper function to detect if the given struct is an error
// returned or not.
func IsErrorStruct(structType interface{}) bool {
	found := false
	errorStruct := []string{
		"BadRequestResponse",
		"UnauthorizedResponse",
		"ErrorRequest",
		"InvalidFormatResponse",
		"AlreadyExistsResponse",
		"NotFoundResponse",
		"ServerErrorResponse",
	}
	st := reflect.TypeOf(structType).Name()

	for _, name := range errorStruct {
		if name == st {
			found = true
			break
		}
	}
	return found
}
