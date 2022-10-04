package ramda

func Tap[T any](fn func(T) any) func(T) T {
	return func(x T) T {
		fn(x)
		return x
	}
}
