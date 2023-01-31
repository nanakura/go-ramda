package ramda

// TODO ChainBy
func chain[T1, T2 any](fn func(T1) []T2) func([]T1) []T2 {
	return func(list []T1) []T2 {
		var result []T2
		for _, it := range list {
			result = append(result, fn(it)...)
		}
		return result
	}
}
