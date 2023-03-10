package ramda

func Scan[T1, T2 any](fn func(T1, T2) T1) func(T1) func([]T2) []T1 {
	return func(acc T1) func([]T2) []T1 {
		return func(list []T2) []T1 {
			length := len(list)
			result := make([]T1, 0, length+1)
			result = append(result, acc)
			for i := 0; i < length; i++ {
				acc = fn(acc, list[i])
				result = append(result, acc)
			}
			return result
		}
	}
}
