package pkg

func MapKeys[T comparable, P any](m map[T]P) []T {
	ks := make([]T, 0, len(m))
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}
