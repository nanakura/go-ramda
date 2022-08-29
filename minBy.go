package ramda

func MinBy[T Numeric](f func(T) T) func(T) func(T) T {
	return func(a T) func(T) T {
		return func(b T) T {
			bb := f(b)
			if Min(f(a))(bb) == bb {
				return b
			}
			return a
		}
	}
}
