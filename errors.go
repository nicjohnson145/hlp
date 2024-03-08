package hlp

func must(err error) {
	if err == nil {
		return
	}

	panic(err.Error())
}

// Must wraps a function call that returns value and error, and panics if the error is non-nil
func Must[T any](out T, err error) T {
	must(err)
	return out
}
