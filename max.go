package ramda

func Max[T Ordered](a T) func(T) T {
	return func(b T) T {
		if a > b {
			return a
		} else {
			return b
		}
	}
}
