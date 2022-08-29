package ramda

// Dec Decrements its argument.
func Dec[T Numeric](num T) T {
	return num - 1
}
