package ramda

func Drop[T any](count int, list ...T) []T {
	if count <= 0 {
		return list
	}

	if count >= len(list) {
		return make([]T, 0)
	}

	return list[count:]
}
