package hashset

type Hasher interface {
	Hash() string
}

type Set[T any] struct {
	hashFunc func(x T) string

	data map[string]T
}

func New[T any](f func(T) string, vals ...T) *Set[T] {
	s := &Set[T]{
		data:     make(map[string]T),
		hashFunc: f,
	}

	for _, val := range vals {
		s.Add(val)
	}

	return s
}

func (s *Set[T]) Contains(val T) bool {
	_, ok := s.data[s.hashFunc(val)]
	return ok
}

func (s *Set[T]) AsSlice() []T {
	out := make([]T, len(s.data))
	i := 0
	for _, val := range s.data {
		out[i] = val
		i += 1
	}
	return out
}

func (s *Set[T]) Count() int {
	out := len(s.data)
	return out
}

func (s *Set[T]) Add(vals ...T) {
	for _, v := range vals {
		s.data[s.hashFunc(v)] = v
	}
}

func (s *Set[T]) Remove(val T) {
	delete(s.data, s.hashFunc(val))
}

func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	newSet := New(s.hashFunc)

	for _, val := range s.data {
		if other.Contains(val) {
			newSet.Add(val)
		}
	}

	return newSet
}

func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	newSet := New(s.hashFunc)

	for _, val := range s.data {
		if !other.Contains(val) {
			newSet.Add(val)
		}
	}

	return newSet
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	newSet := New(s.hashFunc)

	for _, val := range s.data {
		newSet.Add(val)
	}

	for _, val := range other.data {
		newSet.Add(val)
	}

	return newSet
}

func (s *Set[T]) SymmetricDifference(other *Set[T]) *Set[T] {
	newSet := New(s.hashFunc)

	for _, val := range s.data {
		if !other.Contains(val) {
			newSet.Add(val)
		}
	}

	for _, val := range other.data {
		if !s.Contains(val) {
			newSet.Add(val)
		}
	}

	return newSet
}

func (s *Set[T]) Iter() func(func(int, T) bool) {
	count := 0
	return func(yield func(int, T) bool) {
		for _, val := range s.data {
			if !yield(count, val) {
				return
			}
			count += 1
		}
	}
}
