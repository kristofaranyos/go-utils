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

func Filter[T comparable](s []T, f func(t T) bool) []T {
	var res []T

	for _, e := range s {
		if f(e) {
			res = append(res, e)
		}
	}

	return res
}

func Cast[S any, T any](s []S, f func(s S) T) []T {
	result := make([]T, len(s))

	for i, e := range s {
		result[i] = f(e)
	}

	return result
}

// TODO: add tests
