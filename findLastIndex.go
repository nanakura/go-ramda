package ramda

func FindLastIndex[T any](fn func(T) bool) func([]T) int {
	return func(list []T) int {
		lens := len(list)
		for idx := lens - 1; idx >= 0; idx-- {
			if fn(list[idx]) {
				return idx
			}
		}
		return -1
	}
}
