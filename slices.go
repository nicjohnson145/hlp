package hlp

import (
	"cmp"
)

// FilterMapErr is like FilterMap, but the callback function can fail. In such a case, a nil slice and the error from
// the callback is returned. This function fails fast; i.e it stops iteration at the first non-nil error
func FilterMapErr[T any, R any](collection []T, callback func(item T, index int) (R, bool, error)) ([]R, error) {
	result := []R{}

	for i, item := range collection {
		r, ok, err := callback(item, i)
		if err != nil {
			return nil, err
		}
		if !ok {
			continue
		}
		result = append(result, r)
	}

	return result, nil
}

// FilterMap is the combination of Filter & Map, returning the elements from the input slice as transformed by the
// callback function, but only in cases where the callback function also returns true
func FilterMap[T any, R any](collection []T, callback func(item T, index int) (R, bool)) []R {
	result, _ := FilterMapErr[T, R](collection, func(item T, index int) (R, bool, error) {
		out, ok := callback(item, index)
		return out, ok, nil
	})
	return result
}

// MapErr is like Map, but the callback function can fail. In such a case, a nil slice and error from the callback is
// returned. This function fails fast; i.e it stops iteration at the first non-nil error
func MapErr[T any, R any](collection []T, callback func(item T, index int) (R, error)) ([]R, error) {
	return FilterMapErr[T, R](collection, func(item T, index int) (R, bool, error) {
		out, err := callback(item, index)
		return out, true, err
	})
}

// Map returns all elements of the input slice as transformed by the supplied callback function
func Map[T any, R any](collection []T, callback func(item T, index int) R) []R {
	out, _ := FilterMapErr[T, R](collection, func(item T, index int) (R, bool, error) {
		return callback(item, index), true, nil
	})
	return out
}

// FilterErr is like Filter, but the callback function can fail. In such a case, a nil slice and the error from the
// callback is returned. This function fails fast; i.e it stops iteration at the first non-nil error
func FilterErr[T any](collection []T, callback func(item T, index int) (bool, error)) ([]T, error) {
	return FilterMapErr[T, T](collection, func(item T, index int) (T, bool, error) {
		ok, err := callback(item, index)
		return item, ok, err
	})
}

// Filter returns all elements of the input slice the supplied callback returns true for
func Filter[T any](collection []T, callback func(item T, index int) bool) []T {
	out, _ := FilterMapErr[T, T](collection, func(item T, index int) (T, bool, error) {
		return item, callback(item, index), nil
	})
	return out
}

// Flatten consolidates multiple lists into a single conjoined list
func Flatten[T any](lists ...[]T) []T {
	combinedLength := 0

	for _, l := range lists {
		combinedLength += len(l)
	}

	out := make([]T, 0, combinedLength)
	for _, l := range lists {
		out = append(out, l...)
	}

	return out
}

// FillFunc creates an array of length `count` using the return value of the supplied generation function
func FillFunc[T any](count int, genFunc func(int) T) []T {
	out := make([]T, count)
	for i := 0; i < count; i++ {
		out[i] = genFunc(i)
	}

	return out
}

// Fill creates an array of length `count` filled with the supplied value
func Fill[T any](count int, val T) []T {
	return FillFunc(count, func(i int) T { return val })
}

// Batch "chunks" an array into the arrays of the specified size
// Shamelessly stolen from https://go.dev/wiki/SliceTricks
func Batch[T any](iter []T, chunk int) [][]T {
	batches := make([][]T, 0, (len(iter) + chunk - 1) / chunk)
	for chunk < len(iter) {
		iter, batches = iter[chunk:], append(batches, iter[0:chunk:chunk])
	}
	batches = append(batches, iter)

	return batches
}

func GroupBy[T any, U comparable](iter []T, keyGen func(item T) U) map[U][]T {
	result := map[U][]T{}

	for i := range iter {
		key := keyGen(iter[i])

		result[key] = append(result[key], iter[i])
	}

	return result
}

// ExtractRange creates a new list with the elements from list, as specified by the range expression
func ExtractRange[T any](list []T, expr string) ([]T, error) {
	indexList, err := ParseRange(len(list), expr)
	if err != nil {
		return nil, err
	}

	newList := make([]T, len(indexList))
	for i, idx := range indexList {
		newList[i] = list[idx]
	}

	return newList, nil
}

// Any returns true if any element in the list matches the filter function, and false otherwise. Any will stop checking
// after the first success
func Any[T any](list []T, filter func(x T) bool) bool {
	for _, elem := range list {
		if filter(elem) {
			return true
		}
	}
	return false
}

// All returns true if all elements in the list match the filter function, and false otherwise. All will stop checking
// after the first failure
func All[T any](list []T, filter func(x T) bool) bool {
	for _, elem := range list {
		if !filter(elem) {
			return false
		}
	}
	return true
}

// First finds the first element in the list that matchees the filter, returning its index, or -1 on when no filters
// match
func First[T any](list []T, filter func(x T) bool) int {
	for i, elem := range list {
		if filter(elem) {
			return i
		}
	}
	return -1
}

// MaxBy finds the max value in the slice using the given comparison function
func MaxBy[T any](list []T, compare func(item T, higest T) bool) T {
	return compareBy(list, compare)
}

// Max fins the max value in the slice by using the `>` operator
func Max[T cmp.Ordered](list []T) T {
	return compareBy(list, func(a, b T) bool {
		return a > b
	})
}

// MinBy finds the min value in the slice using the given comparison function
func MinBy[T any](list []T, compare func(item T, lowest T) bool) T {
	return compareBy(list, compare)
}

// Min finds the min value in the slice using the `<` operator
func Min[T cmp.Ordered](list []T) T {
	return compareBy(list, func(a, b T) bool {
		return a < b
	})
}

func compareBy[T any](list []T, compare func(a T, b T) bool) T {
	var out T

	if len(list) == 0 {
		return out
	}

	out = list[0]
	for _, item := range list {
		if compare(item, out) {
			out = item
		}
	}

	return out
}
