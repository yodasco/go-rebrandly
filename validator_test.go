package rebrandly

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validDestination(t *testing.T) {
	dest := ""
	err := validDestination(dest)
	assert.EqualError(t, err, FilterErrorEmptyDestination)

	dest = strings.Repeat("0", MaxDestinationLength+1)
	err = validDestination(dest)
	assert.EqualError(t, err, fmt.Sprintf(FilterErrorDestinationBiggerThenMaxLen,
		MaxDestinationLength))

	dest = "%!"
	err = validDestination(dest)
	assert.NotEmpty(t, err)

	dest = "#anchor"
	err = validDestination(dest)
	assert.Empty(t, err)
}
