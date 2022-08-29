package ramda

func Find[T any](fn func(T) bool) func([]T) any {
	return func(list []T) any {
		lens := len(list)
		for idx := 0; idx < lens; idx++ {
			if fn(list[idx]) {
				return list[idx]
			}
		}
		return nil
	}
}
