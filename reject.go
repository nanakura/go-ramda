package ramda

func Reject[T any](fn func(T) bool) func([]T) []T {
	return func(input []T) []T {
		return Filter(func(val T) bool {
			return !fn(val)
		})(input)
	}
}

func RejectWithIndex[T any](fn func(T, int) bool, input []T) func([]T) []T {
	return func(input []T) []T {
		return FilterWithIndex(func(val T, i int) bool {
			return !fn(val, i)
		})(input)
	}
}
