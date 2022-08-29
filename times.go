package ramda

func Times[R any](fn func(int) R) func(int) []R {
	return func(n int) []R {
		list := make([]R, n)
		for i := 0; i < n; i++ {
			list[i] = fn(i)
		}
		return list
	}
}
