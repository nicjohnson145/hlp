package hlp

// MapFromSliceErr is like MapFromSlice, but the callback function can fail. In such a case, a nil slice and the error from
// the callback is returned. This function fails fast; i.e it stops iteration at the first non-nil error
func MapFromSliceErr[T any, K comparable, V any](collection []T, callback func(item T, index int) (K, V, error)) (map[K]V, error) {
	out := map[K]V{}

	for i, item := range collection {
		key, val, err := callback(item, i)
		if err != nil {
			return nil, err
		}
		out[key] = val
	}

	return out, nil
}

// MapFromSlice returns a map, whose keys & values are the return values from applying the callback function to each
// element of the given slice
func MapFromSlice[T any, K comparable, V any](collection []T, callback func(item T, index int) (K, V)) map[K]V {
	out, _ := MapFromSliceErr[T, K, V](collection, func(item T, index int) (K, V, error) {
		k, v := callback(item, index)
		return k, v, nil
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
