package hlp

// Ternary is returns the first argument if condition is true, otherwise returns the second argument
func Ternary[T any](condition bool, a T, b T) T {
	if condition {
		return a
	}
	return b
}
