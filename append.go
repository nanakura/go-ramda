package ramda

// Append Returns a new list containing the content of the given list.
func Append[T any](x ...T) func([]T) []T {
	return func(list []T) []T {
		return append(list, x...)
	}
}
