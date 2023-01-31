package ramda

func DropLast[T any](count int) func([]T) []T {
	return func(list []T) []T {
		listLen := len(list)

		if listLen == 0 || count >= listLen {
			return make([]T, 0)
		}

		res := make([]T, listLen-count)
		copy(res, list[:(listLen-count)])
		return res
	}
}
