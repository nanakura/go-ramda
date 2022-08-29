package ramda

// Add add two values
func Add[T Numeric](a T) func(b T) T {
	return func(b T) T {
		return a + b
	}
}
