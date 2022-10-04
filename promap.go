package ramda

func Promap[T1, T2, T3, T4 any](f func(T1) T2) func(func(T3) T4) func(func(T2) T3) func(T1) T4 {
	return func(g func(T3) T4) func(func(T2) T3) func(T1) T4 {
		return func(h func(T2) T3) func(T1) T4 {
			return func(x T1) T4 {
				return g(h(f(x)))
			}
		}
	}
}
