package ramda

// Values returns a slice of map's values
func Values[T comparable, R any](m map[T]R) []R {
	keys := make([]R, 0, len(m))
	for _, v := range m {
		keys = append(keys, v)
	}
	return keys
}
