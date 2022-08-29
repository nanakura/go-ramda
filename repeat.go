package ramda

func Repeat[T any](value T) func(int) []T {
	return func(n int) []T {
		if n <= 0 {
			return []T{}
		} else {
			list := make([]T, n)
			for i := 0; i < n; i++ {
				list[i] = value
			}
			return list
		}
	}
}
