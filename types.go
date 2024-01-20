package hlp

// Ptr returns a pointer copy of the given value
func Ptr[T any](x T) *T {
	return &x
}

// ToAnySlice returns a new slice where all elements of `collection` are mapped to the `any` type
func ToAnySlice[T any](collection []T) []any {
	return Map[T, any](collection, func(item T, _ int) any {
		return item
	})
}
