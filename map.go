package ramda

// Map Map the values to the function from left to right
func Map[T any, R any](fn func(T) R) func(...T) []R {
	return func(values ...T) []R {
		result := make([]R, len(values))
		for i, val := range values {
			result[i] = fn(val)
		}

		return result
	}
}

// MapIndexed Map the values to the function from left to right
func MapIndexed[T any, R any](fn func(T, int) R) func(...T) []R {
	return func(values ...T) []R {
		result := make([]R, len(values))
		for i, val := range values {
			result[i] = fn(val, i)
		}

		return result
	}
}
