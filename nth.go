package ramda

func Nth[T any](offset int) func([]T) T {
	return func(list []T) T {
		return list[offset]
	}
}
