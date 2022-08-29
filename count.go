package ramda

func Count[T any](pred func(a T) bool) func([]T) int {
	return func(xs []T) int {
		return Reduce(func(a int, e T) int {
			if pred(e) {
				return a + 1
			}
			return a
		})(0)(xs)
	}
}
