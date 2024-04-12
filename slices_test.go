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
