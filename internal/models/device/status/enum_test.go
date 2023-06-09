package status

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	type testCase struct {
		str           string
		expected      Enum
		expectedError error
	}

	tests := []testCase{
		{
			str:      "REACHABLE",
			expected: Reachable,
		},
		{
			str:      "UNREACHABLE",
			expected: Unreachable,
		},
		{
			str:      "MAINTENANCE",
			expected: Maintenance,
		},
		{
			str:           "x_MAINTENANCE_x",
			expectedError: ErrorCouldNotParse,
		},
	}

	for _, test := range tests {
		actual, err := Parse(test.str)
		assert.Equal(t, test.expected, actual)
		assert.ErrorIs(t, test.expectedError, err)
	}
}
