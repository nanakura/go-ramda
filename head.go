package ramda

func Head[T any](list []T) T {
	var result T

	if len(list) <= 0 {
		return result
	}

	result = list[:1][0]

	return result
}
