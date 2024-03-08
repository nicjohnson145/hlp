package hlp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMust(t *testing.T) {
	f := func(x bool) (int, error) {
		if x {
			return 1, nil
		}
		return 0, fmt.Errorf("ahhh")
	}

	t.Run("no panic", func(t *testing.T) {
		require.Equal(t, Must(f(true)), 1)
	})

	t.Run("panic", func(t *testing.T) {
		require.Panics(t, func() { Must(f(false)) })
	})
}
