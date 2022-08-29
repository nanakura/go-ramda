package ramda

func Reject[T any](fn func(T, int) bool, input []T) []T {
	return Filter(func(val T, i int) bool {
		return !fn(val, i)
	})(input)
}
