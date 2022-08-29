package ramda

func Equals[T comparable](a T) func(T) bool {
	return func(b T) bool {
		return a == b
	}
}
