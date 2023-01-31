package ramda

func ForEach[T any](fn func(T)) func([]T) {
	return func(list []T) {
		for _, it := range list {
			fn(it)
		}
	}
}

func ForEachWithIndex[T any](fn func(T, int)) func([]T) {
	return func(list []T) {
		for i, it := range list {
			fn(it, i)
		}
	}
}
