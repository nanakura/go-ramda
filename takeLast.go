package ramda

func TakeLast[T any](count int) func([]T) []T {
	return func(list []T) []T {
		length := len(list)
		if count >= length || count < 0 {
			return list
		} else {
			return list[(length - count):]
		}
	}
}
