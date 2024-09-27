package hlp

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestTernary(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		require.Equal(t, Ternary(true, "a", "b"), "a")
	})

	t.Run("false", func(t *testing.T) {
		require.Equal(t, Ternary(false, "a", "b"), "b")
	})
}
