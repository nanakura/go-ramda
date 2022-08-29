package ramda

func Min[T Ordered](a T) func(T) T {
	return func(b T) T {
		if a > b {
			return b
		} else {
			return a
		}
	}
}
