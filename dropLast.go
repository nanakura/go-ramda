package ramda

func DropLast[T any](count int, list ...T) []T {
	listLen := len(list)

	if listLen == 0 || count >= listLen {
		return make([]T, 0)
	}

	return list[:(listLen - count)]
}
