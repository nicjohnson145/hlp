package hashset

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSet(t *testing.T) {
	intFunc := func(x int) string {
		return fmt.Sprint(x)
	}
	strFunc := func(x string) string {
		return x
	}

	t.Run("contains", func(t *testing.T) {
		s := New(intFunc, 1, 2, 3)
		require.True(t, s.Contains(1))
		require.False(t, s.Contains(4))
	})

	t.Run("remove", func(t *testing.T) {
		s := New(intFunc, 1, 2, 3)
		s.Remove(1)
		require.False(t, s.Contains(1))
	})

	t.Run("count", func(t *testing.T) {
		s := New(intFunc, 1, 2, 3)
		require.Equal(t, 3, s.Count())
	})

	t.Run("intersection", func(t *testing.T) {
		s := New(intFunc, 1, 2, 3)
		o := New(intFunc, 2, 3, 4)

		require.ElementsMatch(t, s.Intersection(o).AsSlice(), []int{2, 3})
	})

	t.Run("difference", func(t *testing.T) {
		s := New(intFunc, 1, 2, 3)
		o := New(intFunc, 2, 3, 4)

		require.ElementsMatch(t, s.Difference(o).AsSlice(), []int{1})
	})

	t.Run("union", func(t *testing.T) {
		s := New(intFunc, 1, 2, 3)
		o := New(intFunc, 2, 3, 4)

		require.ElementsMatch(t, s.Union(o).AsSlice(), []int{1, 2, 3, 4})
	})

	t.Run("symmetric difference", func(t *testing.T) {
		s := New(intFunc, 1, 2, 3)
		o := New(intFunc, 2, 3, 4)

		require.ElementsMatch(t, s.SymmetricDifference(o).AsSlice(), []int{1, 4})
	})
	
	t.Run("iter", func(t *testing.T) {
		s := New(strFunc, "a", "b", "c")
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
