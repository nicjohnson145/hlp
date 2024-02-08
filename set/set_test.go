package set

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSet(t *testing.T) {
	t.Run("contains", func(t *testing.T) {
		s := New(1, 2, 3)
		require.True(t, s.Contains(1))
		require.False(t, s.Contains(4))
	})

	t.Run("remove", func(t *testing.T) {
		s := New(1, 2, 3)
		s.Remove(1)
		require.False(t, s.Contains(1))
	})

	t.Run("count", func(t *testing.T) {
		s := New(1, 2, 3)
		require.Equal(t, 3, s.Count())
	})

	t.Run("intersection", func(t *testing.T) {
		s := New(1, 2, 3)
		o := New(2, 3, 4)

		require.ElementsMatch(t, s.Intersection(o).AsSlice(), []int{2, 3})
	})

	t.Run("difference", func(t *testing.T) {
		s := New(1, 2, 3)
		o := New(2, 3, 4)

		require.ElementsMatch(t, s.Difference(o).AsSlice(), []int{1})
	})

	t.Run("union", func(t *testing.T) {
		s := New(1, 2, 3)
		o := New(2, 3, 4)

		require.ElementsMatch(t, s.Union(o).AsSlice(), []int{1, 2, 3, 4})
	})

	t.Run("symmetric difference", func(t *testing.T) {
		s := New(1, 2, 3)
		o := New(2, 3, 4)

		require.ElementsMatch(t, s.SymmetricDifference(o).AsSlice(), []int{1, 4})
	})
}
