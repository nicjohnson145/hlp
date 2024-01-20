package hlp

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	t.Run("smokes", func(t *testing.T) {
		x := 2
		require.Equal(t, &x, Ptr(2))
	})
}

func TestToAnySlice(t *testing.T) {
	t.Run("smokes", func(t *testing.T) {
		require.Equal(t, []any{1, 2, 3}, ToAnySlice([]int{1, 2, 3}))
	})
}
