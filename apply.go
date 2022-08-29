package ramda

// Apply Applies function `fn` to the argument list `args`
func Apply[T any, R any](fn func(...T) R) func([]T) R {
	return func(args []T) R {
		return fn(args...)
	}
}
