package ramda

func Thunkify[R any](fn func(...any) R) func(...any) func() R {
	return func(args ...any) func() R {
		return func() R {
			return fn(args...)
		}
	}
}
