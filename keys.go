package ramda

// Keys returns a slice of map's keys
func Keys[T comparable, R any](m map[T]R) []T {
	keys := make([]T, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
