package ramda

func O[T1, T2, R any](f func(T2) R) func(func(T1) T2) func(T1) R {
	return func(g func(T1) T2) func(T1) R {
		return func(x T1) R {
			return f(g(x))
		}
	}
}
