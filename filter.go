package ramda

// Filter Filter the values by the given predicate function (predicate func, slice)
func Filter[T any](fn func(T) bool) func([]T) []T {
	return func(input []T) []T {
		list := make([]T, len(input))

		newLen := 0

		for i := range input {
			if fn(input[i]) {
				newLen++
				list[newLen-1] = input[i]
			}
		}

		result := make([]T, newLen)
		copy(result, list[:newLen])
		return result
	}
}

func FilterWithIndex[T any](fn func(T, int) bool) func([]T) []T {
	return func(input []T) []T {
		list := make([]T, len(input))

		newLen := 0

		for i := range input {
			if fn(input[i], i) {
				newLen++
				list[newLen-1] = input[i]
			}
		}

		result := make([]T, newLen)
		copy(result, list[:newLen])
		return result
	}
}
