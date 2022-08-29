package ramda

// Keys returns a slice of map's keys
func Keys[T comparable, R any](m map[T]R) []T {
	keys := make([]T, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}
