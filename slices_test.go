package hlp

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	errInternalTestingError = errors.New("internal testing error")
)

func intSlice() []int {
	return []int{1, 2, 3, 4, 5, 6}
}

func TestFilter(t *testing.T) {
	t.Run("smokes", func(t *testing.T) {
		require.Equal(
			t,
			[]int{2, 4, 6},
			Filter(
				intSlice(),
				func(item, index int) bool {
					return item%2 == 0
				},
			),
		)
	})
}

func TestFilterErr(t *testing.T) {
	t.Run("error case", func(t *testing.T) {
		got, err := FilterErr(
			intSlice(),
			func(item, index int) (bool, error) {
				if item == 3 {
					return false, errInternalTestingError
				}
				return true, nil
			},
		)

		require.Nil(t, got)
		require.ErrorIs(t, err, errInternalTestingError)
	})

	t.Run("non error case", func(t *testing.T) {
		got, err := FilterErr(
			intSlice(),
			func(item, index int) (bool, error) {
				return item % 2 == 0, nil
			},
		)

		require.Equal(t, []int{2, 4, 6}, got)
		require.NoError(t, err)
	})
}

func TestMap(t *testing.T) {
	t.Run("smokes", func(t *testing.T) {
		require.Equal(
			t,
			[]int{2, 4, 6, 8, 10, 12},
			Map(
				intSlice(),
				func(item, index int) int {
					return item * 2
				},
			),
		)
	})
}

func TestMapErr(t *testing.T) {
	t.Run("error case", func(t *testing.T) {
		got, err := MapErr(
			intSlice(),
			func(item, index int) (int, error) {
				if item == 3 {
					return 0, errInternalTestingError
				}
				return item * 2, nil
			},
		)

		require.Nil(t, got)
		require.ErrorIs(t, err, errInternalTestingError)
	})

	t.Run("non error case", func(t *testing.T) {
		got, err := MapErr(
			intSlice(),
			func(item, index int) (int, error) {
				return item * 2, nil
			},
		)

		require.Equal(t, []int{2, 4, 6, 8, 10, 12}, got)
		require.NoError(t, err)
	})
}

func TestFilterMap(t *testing.T) {
	t.Run("smokes", func(t *testing.T) {
		require.Equal(
			t,
			[]int{4, 8, 12},
			FilterMap(intSlice(), func(item, index int) (int, bool) {
				return item * 2, item % 2 == 0
			}),
		)
	})
}

func TestFilterMapErr(t *testing.T) {
	t.Run("error case", func(t *testing.T) {
		got, err := FilterMapErr(
			intSlice(),
			func(item, index int) (int, bool, error) {
				if item == 3 {
					return 0, false, errInternalTestingError
				}
				return item * 2, item % 2 == 0, nil
			},
		)

		require.Nil(t, got)
		require.ErrorIs(t, err, errInternalTestingError)
	})

	t.Run("non error case", func(t *testing.T) {
		got, err := FilterMapErr(
			intSlice(),
			func(item, index int) (int, bool, error) {
				return item * 2, item % 2 == 0, nil
			},
		)

		require.Equal(t, []int{4, 8, 12}, got)
		require.NoError(t, err)
	})
}

func TestFlatten(t *testing.T) {
	t.Run("smokes", func(t *testing.T) {
		result := Flatten(
			[]string{"a", "b", "c"},
			[]string{"d", "e", "f"},
		)

		require.Equal(t, []string{"a", "b", "c", "d", "e", "f"}, result)
	})
}

func TestFillFunc(t *testing.T) {
	t.Run("smokes", func(t *testing.T) {
		results := FillFunc(3, func(i int) int { return i * 3 })
		require.Equal(t, []int{0, 3, 6}, results)
	})
}

func TestFill(t *testing.T) {
	t.Run("smokes", func(t *testing.T) {
		results := Fill(3, "abc")
		require.Equal(t, []string{"abc", "abc", "abc"}, results)
	})
}

func TestBatch(t *testing.T) {
	t.Run("smokes", func(t *testing.T) {
		results := Batch([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 3)
		require.Equal(
			t,
			[][]int{
				{0, 1, 2},
				{3, 4, 5},
				{6, 7, 8},
				{9},
			},
			results,
		)
	})
}

func TestGroupBy(t *testing.T) {
	t.Run("smokes", func(t *testing.T) {
		result := GroupBy([]string{"apple", "amazon", "bravo", "bakery", "cherry", "chocolate"}, func(item string) string {
			return string(item[0])
		})
		require.Equal(
			t,
			map[string][]string{
				"a": {"apple", "amazon"},
				"b": {"bravo", "bakery"},
				"c": {"cherry", "chocolate"},
			},
			result,
		)
	})
}
