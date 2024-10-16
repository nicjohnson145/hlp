package hlp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDefaultEnv(t *testing.T) {
	t.Run("set", func(t *testing.T) {
		t.Setenv("FOO", "BAR")
		require.Equal(t, DefaultEnv("FOO", "BAZ"), "BAR")
	})

	t.Run("not set", func(t *testing.T) {
		require.Equal(t, DefaultEnv("FOO", "BAZ"), "BAZ")
	})
}
