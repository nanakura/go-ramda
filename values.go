package ramda

// Values returns a slice of map's values
func Values[T comparable, R any](m map[T]R) []R {
	keys := make([]R, len(m))
	i := 0
	for _, v := range m {
		keys[i] = v
		i++
	}
	return keys
}
