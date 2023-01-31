package ramda

func Take[T any](count int) func([]T) []T {
	return func(list []T) []T {
		if count >= len(list) || count < 0 {
			return list
		} else {
			res := make([]T, count)
			copy(res, list[:count])
			return res
		}
	}
}
