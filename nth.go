package ramda

func Nth[T any](offset int) func([]T) T {
	return func(list []T) T {
		if offset < 0 {
			offset = len(list) + offset
		}
		return list[offset]
	}
}
