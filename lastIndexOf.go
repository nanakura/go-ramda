package ramda

func LastIndexOf[T comparable](target T) func([]T) int {
	return func(xs []T) int {
		lens := len(xs)
		for idx := lens - 1; idx >= 0; idx-- {
			if xs[idx] == target {
				return idx
			}
		}
		return -1
	}
}
