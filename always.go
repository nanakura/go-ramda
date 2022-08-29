package ramda

// Always Returns a function that always returns the given value.
func Always[T any](value T) func() T {
	return func() T {
		return value
	}
}
