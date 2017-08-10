package rebrandly

// ErrorCode holds a machine readable status for the error
type ErrorCode string

// Error codes that are returned back
const (
	ErrorCodeUnauthorized        ErrorCode = "Unauthorized"
	ErrorCodeInvalidFormat       ErrorCode = "InvalidFormat"
	ErrorCodeRequiredField       ErrorCode = "RequiredField"
	ErrorCodeInvalidLength       ErrorCode = "InvalidLength"
	ErrorCodeInvalidMinLength    ErrorCode = "InvalidMinLength"
	ErrorCodeInvalidMaxLength    ErrorCode = "InvalidMaxLength"
	ErrorCodeInvalidEmailAddress ErrorCode = "InvalidEmailAddress"
	ErrorCodeOutOfRange          ErrorCode = "OutOfRange"
	ErrorCodePatternMismatch     ErrorCode = "PatternMismatch"
	ErrorCodePrefixMismatch      ErrorCode = "PrefixMismatch"
	ErrorCodeInvalidCharacter    ErrorCode = "InvalidCharacter"
	ErrorCodeMustBeLowerCase     ErrorCode = "MustBeLowerCase"
	ErrorCodeMustBeUpperCase     ErrorCode = "MustBeUpperCase"
	ErrorCodeAlreadyExists       ErrorCode = "AlreadyExists"
	ErrorCodeNotFound            ErrorCode = "NotFound"
)

// BadRequestResponse is the responce when a given JSON structure is invalid.
//
// Example JSON error
//   {
//      "Message": "Unexpected token a"
//   }
type BadRequestResponse struct {
	// Message to the user explaining what is wrong
	Message string `json:"message"`
}

// UnauthorizedResponse is a struct for holding Unauthorized responce from the
// server
//
// Example authorization error
//   {
//     "code": "Unauthorized",
//     "message": "Missing OAuth token in request"
//   }
type UnauthorizedResponse struct {
	// A message to the user explaining why the request was not authorized
	Message string `json:"message"`
	// Always "Unauthorized"
	Code ErrorCode `json:"code"`
}

// ErrorRequest holds a structure for nested errors
type ErrorRequest struct {
	// Message to the user explaining what happened
	Message string `json:"message"`
	// Machine readable code to handle the error
	Code ErrorCode `json:"code"`
	// Request property which originated the error
	Property string `json:"property"`
}

// InvalidFormatResponse is a struct that contains details regarding why a given
// request had invalid format.
//
// Example validation error: missing required field
//   {
//     "code": "RequiredField",
//     "message": "Cannot be empty",
//     "property": "slashtag"
//   }
//
// Example of validation error: invalid min length
//   {
//     "property": "slashtag",
//     "message": "Value cannot be less than 2 characters long",
//     "code": "InvalidMinLength",
//     "input": "a",
//     "minLength": 2
//   }
//
// Example of validation error: OutOfRange
//   {
//     "property": "regType",
//     "message": "Value is not allowed",
//     "code": "OutOfRange",
//     "input": "cat",
//     "range": ["individual", "business"]
//   }
type InvalidFormatResponse struct {
	// Message to the user explaining what happened
	Message string `json:"message"`
	// Machine readable code to handle the error
	Code ErrorCode `json:"code"`
	// Request property which originated the error
	Property string `json:"property"`
	// Message to the developer further explaining what happened
	Verbose string `json:"verbose"`
	// Original input value for property
	Input interface{} `json:"input"`
	// Minimum length allowed
	MinLength uint64 `json:"minLength"`
	// Maximum length allowed
	MaxLength uint64 `json:"maxLength"`
	// Length to match
	Length uint64 `json:"length"`
	// Set of allowed values
	Range []interface{} `json:"range"`
	// Regex to test the input with
	Pattern string `json:"pattern"`
	// Prefix that input has to match
	Prefix string `json:"prefix"`
	// Character given in input
	Character string `json:"character"`
	// list of errors
	Errors []ErrorRequest `json:"errors"`
}

// AlreadyExistsResponse indicates that it is not possible to create a resource
// with the given definition, because another resource already exists with the
// same attributes.
//
// Example Already Exists error
//   {
//     "property": "slashtag",
//     "message": "Already exists",
//     "code": "AlreadyExists"
//   }
type AlreadyExistsResponse struct {
	// Always "Already exists"
	Message string `json:"message"`
	// Always "AlreadyExists"
	Code ErrorCode `json:"code"`
	// Request property which originated the error
	Property string `json:"property"`
}

// NotFoundResponse holds information of what was not found
//
// Example Not Found error
//   {
//     "property": "id",
//     "message": "Not found",
//     "code": "NotFound"
//   }
type NotFoundResponse struct {
	// Always "Not found"
	Message string `json:"message"`
	// Always "NotFound"
	Code ErrorCode `json:"code"`
	// Request property which originated the error
	Property string `json:"property"`

	// Undocumented fields

	Source string `json:"source"`
	ID     int64  `json:"id"`
}

// ServerErrorResponse occurs when something went unexpectedly wrong with API
// operation on the server side.
//
// Example generic error
//   {
//     "message": "An error occurred"
//   }
type ServerErrorResponse struct {
	// Message to user explaining what happened
	Message string `json:"message"`
}
