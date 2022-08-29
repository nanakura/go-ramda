package ramda

func FindIndex[T any](fn func(T) bool) func([]T) int {
	return func(list []T) int {
		lens := len(list)
		for idx := 0; idx < lens; idx++ {
			if fn(list[idx]) {
				return idx
			}
		}
		return -1
	}
}
