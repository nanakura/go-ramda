package ramda

func MaxBy[T Numeric](f func(T) T) func(T) func(T) T {
	return func(a T) func(b T) T {
		return func(b T) T {
			bb := f(b)
			if Max(f(a))(bb) == bb {
				return b
			}
			return a
		}
	}
}
