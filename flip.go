package ramda

func Flip[T any, R any](fn func(T) func(T) R) func(T) func(T) R {
	return func(a T) func(T) R {
		return func(b T) R {
			return fn(b)(a)
		}
	}
}
