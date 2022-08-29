package ramda

func ForEach[T any](fn func(T, int)) func([]T) {
	return func(list []T) {
		for i, it := range list {
			fn(it, i)
		}
	}
}
