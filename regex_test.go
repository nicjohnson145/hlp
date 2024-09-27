package hlp

import (
	"regexp"
	"testing"
	"github.com/stretchr/testify/require"
)

func TestExtractNamedMatches(t *testing.T) {
	t.Run("smokes", func(t *testing.T) {
		exp := regexp.MustCompile(`(?P<first>\d+)\.(?P<second>\d+)\.(\d+)`)
		got := ExtractNamedMatches(exp, exp.FindStringSubmatch("192.168.1"))
		require.Equal(
			t,
			map[string]string{
				"first": "192",
				"second": "168",
			},
			got,
		)
	})
}
