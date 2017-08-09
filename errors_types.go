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

// BasicErrorResponse is a struct for common properties for all errors (except
// server error)
type BasicErrorResponse struct {
	Message string    `json:"message"`
	Code    ErrorCode `json:"code"`
}

// UnauthorizedResponse is a struct for holding Unauthorized responce from the
// server
// Message - A message to the user explaining why the request was not authorized
// Code - Always "Unauthorized"
type UnauthorizedResponse struct {
	BasicErrorResponse
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
	BasicErrorResponse

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
	BasicErrorResponse

	// Request property which originated the error
	Property string `json:"property"`
}
