package ramda

// And returns the first argument if it is falsy, otherwise the second argument.
func And(a bool) func(bool) bool {
	return func(b bool) bool {
		return a && b
	}
}
