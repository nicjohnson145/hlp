//go:build goexperiment.rangefunc
package set

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestIter(t *testing.T) {
	t.Run("smokes", func(t *testing.T) {
		s := New("a", "b", "c")
		ids := []int{}
		vals := []string{}
		for i, val := range s.Iter() {
			ids = append(ids, i)
			vals = append(vals, val)
		}

		require.Equal(t, []int{0, 1, 2}, ids)
		require.ElementsMatch(t, []string{"a", "b", "c"}, vals)
	})
}
