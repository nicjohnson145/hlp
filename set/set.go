package set

func New[T comparable](vals ...T) *Set[T] {
	s := &Set[T]{
		data: make(map[T]struct{}),
	}
	if len(vals) > 0 {
		s.Add(vals...)
	}
	return s
}

type Set[T comparable] struct {
	data map[T]struct{}
}

func (s *Set[T]) AsSlice() []T {
	out := []T{}
	for key := range s.data {
		out = append(out, key)
	}
	return out
}

func (s *Set[T]) Count() int {
	return len(s.data)
}

func (s *Set[T]) Add(vals ...T) {
	for _, v := range vals {
		v := v
		s.data[v] = struct{}{}
	}
}

func (s *Set[T]) Contains(val T) bool {
	_, ok := s.data[val]
	return ok
}

func (s *Set[T]) Remove(val T) {
	delete(s.data, val)
}

func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	newSet := New[T]()

	for key := range s.data {
		if other.Contains(key) {
			newSet.Add(key)
		}
	}

	return newSet
}

func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	newSet := New[T]()

	for key := range s.data {
		if !other.Contains(key) {
			newSet.Add(key)
		}
	}

	return newSet
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	newSet := New[T]()

	for key := range s.data {
		newSet.Add(key)
	}

	for key := range other.data {
		newSet.Add(key)
	}

	return newSet
}

func (s *Set[T]) SymmetricDifference(other *Set[T]) *Set[T] {
	newSet := New[T]()

	for key := range s.data {
		if !other.Contains(key) {
			newSet.Add(key)
		}
	}

	for key := range other.data {
		if !s.Contains(key) {
			newSet.Add(key)
		}
	}

	return newSet
}

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
