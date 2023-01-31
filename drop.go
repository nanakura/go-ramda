package ramda

func Drop[T any](count int) func([]T) []T {
	return func(list []T) []T {
		if count <= 0 {
			return list
		}
		listLen := len(list)
		if count >= listLen {
			return make([]T, 0)
		}
		res := make([]T, listLen-count)
		copy(res, list[count:])
		return res
	}
}
