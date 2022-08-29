package ramda

func FindLast[T any](fn func(T) bool) func([]T) any {
	return func(list []T) any {
		lens := len(list)
		for idx := lens - 1; idx >= 0; idx-- {
			if fn(list[idx]) {
				return list[idx]
			}
		}
		return nil
	}
}
