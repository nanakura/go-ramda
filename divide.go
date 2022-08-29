package ramda

func Divide[T Numeric](a T) func(T) T {
	return func(b T) T {
		return a / b
	}
}
