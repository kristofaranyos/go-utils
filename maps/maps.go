package maps

func Contains[K comparable, V comparable](m map[K]V, v V) bool {
	for _, e := range m {
		if e == v {
			return true
		}
	}

	return false
}

func ContainsKey[K comparable, V comparable](m map[K]V, k K) bool {
	for e, _ := range m {
		if e == k {
			return true
		}
	}

	return false
}

func Clear[K comparable, V comparable](m map[K]V) {
	for k, _ := range m {
		delete(m, k)
	}
}
