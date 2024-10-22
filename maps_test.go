package hlp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMapFromSlice(t *testing.T) {
	type thing struct {
		ID int
		Name string
	}

	input := func() []thing{
		return []thing{
			{ID: 1, Name: "First"},
			{ID: 2, Name: "Second"},
		}
	}

	t.Run("base variant", func(t *testing.T) {
		want := map[string]thing{
			"First": {ID: 1, Name: "First"},
			"Second": {ID: 2, Name: "Second"},
		}

		require.Equal(t, want, MapFromSlice(input(), func(item thing, index int) (string, thing) {
			return item.Name, item
		}))
	})

	t.Run("error variant", func(t *testing.T) {
		t.Run("error case", func(t *testing.T) {
			got, err := MapFromSliceErr(input(), func(item thing, index int) (string, thing, error) {
				if index == 1 {
					return "", thing{}, errInternalTestingError
				}
				return item.Name, item, nil
			})
			require.Nil(t, got)
			require.ErrorIs(t, err, errInternalTestingError)
		})

		t.Run("non-error case", func(t *testing.T) {
			want := map[string]thing{
				"First": {ID: 1, Name: "First"},
				"Second": {ID: 2, Name: "Second"},
			}

			got, err := MapFromSliceErr(input(), func(item thing, index int) (string, thing, error) {
				return item.Name, item, nil
			})
			require.NoError(t, err)
			require.Equal(t, want, got)
		})
	})
}

func TestKeys(t *testing.T) {
	t.Run("smokes", func(t *testing.T) {
		m := map[string]int{
			"1": 1,
			"2": 2,
		}

		require.ElementsMatch(t, []string{"1", "2"}, Keys(m))
	})
}

func TestValues(t *testing.T) {
	t.Run("smokes", func(t *testing.T) {
		m := map[string]int{
			"1": 1,
			"2": 2,
		}

		require.ElementsMatch(t, []int{1, 2}, Values(m))
	})
}

func TestAssign(t *testing.T) {
	m1 := map[string]int{
		"A": 1,
		"B": 2,
	}
	m2 := map[string]int{
		"A": 0,
		"C": 3,
	}

	require.Equal(
		t,
		map[string]int{
			"A": 0,
			"B": 2,
			"C": 3,
		},
		Assign(m1, m2),
	)
}


func TestInvert(t *testing.T) {
	t.Run("smokes", func(t *testing.T) {
		require.Equal(
			t,
			map[int]string{1: "a", 2: "b"},
			Invert(map[string]int{"a": 1, "b": 2}),
		)
	})
}
