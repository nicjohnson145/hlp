//go:build goexperiment.rangefunc
package set

func (s *Set[T]) Iter() func(func(int, T) bool) {
	count := 0
	return func(yield func(int, T) bool) {
		for key := range s.data {
			if !yield(count, key) {
				return
			}
			count += 1
		}
	}
}
