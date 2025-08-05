package hlp

// MapFromSliceErr is like MapFromSlice, but the callback function can fail. In such a case, a nil slice and the error from
// the callback is returned. This function fails fast; i.e it stops iteration at the first non-nil error
func MapFromSliceErr[T any, K comparable, V any](collection []T, callback func(item T, index int) (K, V, error)) (map[K]V, error) {
	out, err := FilteredMapFromSliceErr(collection, func(item T, index int) (K, V, bool, error) {
		key, val, err := callback(item, index)
		return key, val, true, err
	})
	return out, err
}

// MapFromSlice returns a map, whose keys & values are the return values from applying the callback function to each
// element of the given slice
func MapFromSlice[T any, K comparable, V any](collection []T, callback func(item T, index int) (K, V)) map[K]V {
	out, _ := MapFromSliceErr(collection, func(item T, index int) (K, V, error) {
		k, v := callback(item, index)
		return k, v, nil
	})
	return out
}

// FilteredMapFromSliceErr returns a map whose keys and values are the return values from callback function to each
// element of the slice, where the callback function returns true. If the callback function returns error, iteration
// stops and the return value is a nil map and the error
func FilteredMapFromSliceErr[T any, K comparable, V any](collection []T, callback func(item T, index int) (K, V, bool, error)) (map[K]V, error) {
	out := map[K]V{}

	for i, item := range collection {
		key, val, ok, err := callback(item, i)
		if err != nil {
			return nil, err
		}
		if !ok {
			continue
		}
		out[key] = val
	}

	return out, nil
}

// FilteredMapFromSlice is the same as FilteredMapFromSliceErr, except the callback cannot return error
func FilteredMapFromSlice[T any, K comparable, V any](collection []T, callback func(item T, index int) (K, V, bool)) map[K]V {
	out, _ := FilteredMapFromSliceErr(collection, func(item T, index int) (K, V, bool, error) {
		key, val, ok := callback(item, index)
		return key, val, ok, nil
	})
	return out
}

// Assign merges multiple maps from left to right.
func Assign[K comparable, V any](maps ...map[K]V) map[K]V {
	out := map[K]V{}

	for _, m := range maps {
		for k, v := range m {
			out[k] = v
		}
	}

	return out
}

// Keys creates an slice of the map keys.
func Keys[K comparable, V any](m map[K]V) []K {
	out := []K{}

	for k := range m {
		out = append(out, k)
	}

	return out
}

// Values creates an slice of the map values.
func Values[K comparable, V any](m map[K]V) []V {
	out := []V{}

	for _, v := range m {
		out = append(out, v)
	}

	return out
}

// Invert swaps keys & values for the given map
func Invert[K comparable, V comparable](m map[K]V) map[V]K {
	out := map[V]K{}

	for key, val := range m {
		out[val] = key
	}

	return out
}

// SliceFromMap is like FilteredSliceFromMap but no filtering takes place, all entries are mapped
func SliceFromMap[K comparable, V any, T any](m map[K]V, callback func(k K, v V) T) []T {
	return FilteredSliceFromMap(m, func(k K, v V) (T, bool) {
		return callback(k, v), true
	})
}

// FilteredSliceFromMap maps the given map into a slice of type T, for all entries where the callback function returns
// true
func FilteredSliceFromMap[K comparable, V any, T any](m map[K]V, callback func(k K, v V) (T, bool)) []T {
	out := []T{}

	for key, val := range m {
		got, ok := callback(key, val)
		if ok {
			out = append(out, got)
		}
	}

	return out
}
