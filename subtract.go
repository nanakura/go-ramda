package ramda

func Subtract[T Numeric](a T) func(T) T {
	return func(b T) T {
		return a - b
	}
}
