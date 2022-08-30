package ramda

func Product[T Numeric](list []T) T {
	return Reduce[T, T](func(a, b T) T {
		return Multiply(a)(b)
	})(1)(list)
}
