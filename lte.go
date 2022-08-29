package ramda

func Lte[T Ordered](a T) func(T) bool {
	return func(b T) bool {
		return a <= b
	}
}
