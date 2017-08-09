package rebrandly

import (
	"errors"
	"fmt"
	"net/url"
)

const (
	//MaxDestinationLength defined on the docs as Max 1000 chars
	MaxDestinationLength = 1000
)

// Error messages for the filters
const (
	FilterErrorEmptyDestination            = "Destination cannot be empty"
	FilterErrorDestinationBiggerThenMaxLen = "Destination length is bigger then %d chars"
)

func validDestination(destination string) error {
	l := len(destination)
	if l == 0 {
		return errors.New(FilterErrorEmptyDestination)
	}

	if l > MaxDestinationLength {
		return fmt.Errorf(FilterErrorDestinationBiggerThenMaxLen,
			MaxDestinationLength)
	}

	_, err := url.Parse(destination)
	if err != nil {
		return err
	}

	return nil
}
