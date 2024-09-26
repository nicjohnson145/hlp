package hlp

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseRange(t *testing.T) {
	testData := []struct {
		name     string
		expr     string
		expected []int
		error    error
	}{
		{
			name:     "single value",
			expr:     "3",
			expected: []int{3},
		},
		{
			name:     "multiple values",
			expr:     "3,4,6",
			expected: []int{3,4,6},
		},
		{
			name:     "one range",
			expr:     "3-5",
			expected: []int{3,4,5},
		},
		{
			name:     "multiple ranges",
			expr:     "2-4,6-8",
			expected: []int{2,3,4,6,7,8},
		},
		{
			name:     "open range",
			expr:     "7-",
			expected: []int{7,8,9},
		},
		{
			name:     "real complex",
			expr:     "2,4-6,8-",
			expected: []int{2,4,5,6,8,9},
		},
		{
			name:     "real complex",
			expr:     "2,4-6,8-",
			expected: []int{2,4,5,6,8,9},
		},
	}
	for _, tc := range testData {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseRange(10, tc.expr)
			if tc.error != nil {
				require.ErrorIs(t, err, tc.error)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, got)
			}
		})
	}
}
