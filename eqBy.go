package ramda

func EqBy[A any, B comparable](f func(A) B) func(A) func(A) bool {
	return func(x A) func(A) bool {
		return func(y A) bool {
			return f(x) == f(y)
		}
	}
}
