package slices

func Contains[T comparable](s []T, t T) bool {
	for _, e := range s {
		if e == t {
			return true
		}
	}

	return false
}

func Map[T comparable](s []T, f func(t T) T) []T {
	res := make([]T, len(s))

	for i, e := range s {
		res[i] = f(e)
	}

	return res
}

// TODO: add tests
