package ramda

func TakeLast[T any](count int) func([]T) []T {
	return func(list []T) []T {
		length := len(list)
		if count >= length || count < 0 {
			return list
		} else {
			res := make([]T, count)
			copy(res, list[length-count:])
			return res
		}
	}
}
